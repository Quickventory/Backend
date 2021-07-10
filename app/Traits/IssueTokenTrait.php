<?php

namespace App\Traits;

use Illuminate\Http\Request;
use App\Models\User;
use Exception;
use Illuminate\Database\Eloquent\ModelNotFoundException;

trait IssueTokenTrait
{
    public function issueToken(Request $request, $grantType, $scope = "")
    {
        if (!isset($this->client) || is_null($this->client)) {
            die("You should try the command 'php artisan passport:client --password' before using the local backend.");
        }

        $params = [
            'grant_type' => $grantType,
            'client_id' => $this->client->id,
            'client_secret' => $this->client->secret,
            'scope' => $scope
        ];

        if ($grantType !== 'refresh_token') {
            $params['username'] = $request->username;
            $params['password'] = $request->password;
            /** @var User $user */
            $user = User::where('email', $request->username)->orWhere('username', $request->username)->first();

            if (!$user) {
                throw new ModelNotFoundException();
            }

            if ($user->last_token) {

                $user->last_refresh_token->update([
                    'revoked' => true
                ]);

                $user->last_token->update([
                    'revoked' => true
                ]);
            }
        }

        if ($grantType === 'refresh_token') {
            $params['refresh_token'] = $request->refresh_token;
        }


        $request->request->add($params);

        $proxy = Request::create('oauth/token', 'POST', $params);

        $result = app()->handle($proxy);

        if (strpos($result->getContent(), 'refresh token is invalid') !== false) {
            throw new Exception("Invalid refresh token");
        }

        if (strpos($result->getContent(), 'user credentials were incorrect') !== false) {
            throw new Exception("Invalid credentials");
        }

        return $result;
    }
}