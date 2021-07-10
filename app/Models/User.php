<?php

namespace App\Models;

use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Illuminate\Notifications\Notifiable;
use Laravel\Passport\HasApiTokens;
use Laravel\Passport\RefreshToken;

/**
 * [Description User]
 * @property-read integer $id 
 * @property string $first_name 
 * @property string $last_name 
 * @property string $email 
 * @property string $password 
 * 
 * @property-read Customer[] $customers 
 */
class User extends Authenticatable
{
    use HasFactory, Notifiable, HasApiTokens;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'first_name',
        'last_name',
        'email',
        'password',
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password',
        'remember_token',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [
        'email_verified_at' => 'datetime',
    ];

    public function getLastTokenAttribute()
    {
        return $this->tokens()->orderBy('created_at', 'DESC')->whereRevoked(false)->first();
    }

    public function getLastRefreshTokenAttribute()
    {
        return RefreshToken::whereAccessTokenId($this->last_token->id)->first();
    }

    public function findForPassport($identifier)
    {
        return $this->where('email', $identifier);
    }

    public function customers() {
        return $this->hasMany(Customer::class);
    }
}
