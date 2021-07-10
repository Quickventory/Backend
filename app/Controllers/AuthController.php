<?php

namespace App\Controllers;

use App\Models\User;
use Laravel\Passport\Client;
use App\Traits\PassportToken;
use App\Requests\LoginRequest;
use App\Traits\IssueTokenTrait;
use App\Requests\RegisterRequest;
use Illuminate\Support\Facades\DB;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Auth;
use App\Requests\RefreshTokenRequest;
use App\Models\Customer;
use Carbon\Carbon;
use Illuminate\Support\Str;

/**
 * [Description AuthController]
 */
class AuthController extends Controller
{

    use PassportToken, IssueTokenTrait;

    /**
     * @var Client
     */
    private Client $client;
    protected $model = User::class;

    public function __construct()
    {
        $this->client = Client::where('personal_access_client', false)->firstOrFail();
    }

    /**
     * @param RegisterRequest $request
     * 
     * @return [type]
     */
    public function register(RegisterRequest $request)
    {
        try {
            $request->validated();
            /** @var User $user */
            $user = User::create([
                'first_name' => $request->first_name,
                'last_name' => $request->last_name,
                'email' => $request->email,
                'password' => \bcrypt($request->password),
                'date_of_birth' => $request->date_of_birth
            ]);

            $password = $request->automatically_generate_db_password ? Str::random(16) : bcrypt($request->password);
            /** @var Customer $customer */
            $customer = $user->customers()->create([
                'customer_name' => $request->customer_name,
                'tenancy_db_username' => $request->tenancy_db_username,
                'tenancy_db_password' => $request->automatically_generate_db_password ? \bcrypt($password) : bcrypt($request->password),
            ]);

            $customer->domains()->create([
                'domain' => Str::snake($request->customer_name)
            ]);

            $customer->run(function() use($user)  {
               User::create([
                   'first_name' => $user->first_name,
                   'last_name' => $user->last_name,
                   'email' => $user->email,
                   'password' => $user->password
               ]);
            });

        } catch (\Exception $e) {
            DB::rollBack();
            return $this->errorResponse();
        }

        $bearerTokenInfo = $this->getBearerTokenByUser($user->refresh(), $this->client->id);

        return $this->okResponse([
            'bearer' => $bearerTokenInfo,
            'password' => $request->automatically_generate_db_password ? $password : null
        ]);
    }

    /**
     * @param LoginRequest $request
     * 
     * @return [type]
     */
    public function login(LoginRequest $request)
    {
        $validated = $request->validated();
        $username = $validated['username'];
        $password = $validated['password'];
        if (filter_var($username, FILTER_VALIDATE_EMAIL)) {
            Auth::attempt(['email' => $username, 'password' => $password]);
        } else {
            Auth::attempt(['username' => $username, 'password' => $password]);
        }

        if (Auth::check()) {
            return $this->issueToken($request, 'password', '*');
        }

        return $this->errorResponse([
            'message' => 'Invalid credentials'
        ]);
    }

    /**
     * @param RefreshTokenRequest $request
     * 
     * @return [type]
     */
    public function refresh(RefreshTokenRequest $request)
    {
        return $this->issueToken($request, 'refresh_token', '*');
    }
}
