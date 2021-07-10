<?php

namespace App\Models;

use App\Models\Customer;
use Stancl\Tenancy\Database\Models\Tenant;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Stancl\Tenancy\Database\Models\Domain as ModelsDomain;

/**
 * @property-read integer $id
 * @property string $domain
 * @property integer $customer_id
 *
 * @property-read Customer| $customer
 * @property-read Tenant| $tenant
 */
class CustomerDomain extends ModelsDomain
{
    use HasFactory;

    protected $table = 'customer_domains';

    /**
     * @return BelongsTo|Customer
     */
    public function customer() {
        return $this->belongsTo(Customer::class);
    }

    /** 
     * We need this method in order to overwrite the default key on which it tries to fetch the tenant for Cache
     * @return BelongsTo|Tenant
     */
    public function tenant()
    {
        return $this->belongsTo(Customer::class, 'customer_id');
    }
}
