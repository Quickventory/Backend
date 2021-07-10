<?php

namespace App\Requests;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\Rules\Password;

/**
 * Class CustomerCreationPostRequest
 * @package App\Requests
 *
 * @property-read string $first_name
 * @property-read string $last_name
 * @property-read string $customer_name
 * @property-read string $email
 * @property-read Carbon $date_of_birth
 * @property-read string $tenancy_db_username
 * @property-read string $tenancy_db_password
 * @property-read string $password
 * @property-read bool $automatically_generate_db_password
 */
class CustomerCreationPostRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'first_name' => 'required|string',
            'last_name' => 'required|string',
            'customer_name' => 'required|string|max:16',
            'email' => 'required|string|email:rfc,dns|unique:users|max:255',
            'date_of_birth' => 'required|date_format:Y-m-d|',
            'tenancy_db_username' => 'required|string|max:16',
            'tenancy_db_password' => [
                'required_if:automatically_generate_password,false',
                'string',
                'confirmed',
                Password::min(8)
                    ->mixedCase()
                    ->letters()
                    ->numbers()
                    ->symbols()
                    ->uncompromised(),
            ],
            'password' => [
                'required',
                'string',
                'confirmed',
                Password::min(8)
                    ->mixedCase()
                    ->letters()
                    ->numbers()
                    ->symbols()
                    ->uncompromised(),
            ],
            'automatically_generate_db_password' => 'required|bool'
        ];
    }
}
