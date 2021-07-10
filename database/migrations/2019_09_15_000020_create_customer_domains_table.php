<?php

declare(strict_types=1);

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateCustomerDomainsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up(): void
    {
        Schema::create('customer_domains', function (Blueprint $table) {
            $table->id();
            $table->string('domain', 255)->unique();
            $table->unsignedBigInteger('customer_id')->index();
            $table->timestamps();
            $table->foreign('customer_id')->references('id')->on('customers');
            $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down(): void
    {
        Schema::dropIfExists('customer_domains');
    }
}
