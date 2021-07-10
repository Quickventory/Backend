<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Support\Facades\DB;
use Illuminate\Http\Response;

class TransactionMiddleware
{
    /**
     * Handle an incoming request.
     *
     * @param \Illuminate\Http\Request $request
     * @param \Closure $next
     *
     * @return mixed
     * @throws \Throwable
     */
    public function handle($request, Closure $next)
    {
        DB::beginTransaction();

        /** @var Response $response */
        $response = $next($request);

        if ((property_exists($response, 'exception') && $response->exception) || $response->getStatusCode() > 399) {
            while (DB::transactionLevel() > 0) {
                DB::rollBack();
            }
        } else {
            while (DB::transactionLevel() > 0) {
                DB::commit();
            }
        }

        return $response;
    }
}
