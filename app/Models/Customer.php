<?php

namespace App\Models;

use App\Models\CustomerDomain;
use Stancl\Tenancy\Contracts\TenantWithDatabase;
use Stancl\Tenancy\Database\Concerns\HasDomains;
use Stancl\Tenancy\Database\Concerns\HasDatabase;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Stancl\Tenancy\Database\Models\Tenant as BaseTenant;

/**
 * @property-read integer $id
 * @property string $customer_db_name
 * @property integer $user_id
 * 
 * @property-read CustomerDomain $domain
 * 
 */
class Customer extends BaseTenant implements TenantWithDatabase
{
    use HasFactory, HasDatabase, HasDomains;

    protected $table = 'customers';

    public function getTenantKeyName(): string
    {
        return 'customer_name';
    }

    public static function getCustomColumns(): array
    {
        return [
            'id',
            'customer_name',
            'user_id',
            'tenancy_db_username',
            'tenancy_db_password',
            'stripe_id',
            'pm_type',
            'pm_last_four',
            'trial_ends_at',
            'created_at',
            'updated_at',
            'deleted_at',
        ];
    }

    public function getIncrementing()
    {
        return true;
    }

    public function getTenantKey()
    {
        return $this->getAttribute($this->getTenantKeyName());
    }

    public function domains() {
        return $this->hasMany(CustomerDomain::class);
    }
}
