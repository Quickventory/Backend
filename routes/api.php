<?php

use Illuminate\Support\Facades\Route;
use App\Controllers\AuthController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware(['throttle:25,1', 'auth:api', InitializeTenancyBySubdomain::class, PreventAccessFromCentralDomains::class])->group(function () {
    Route::get('/me', [AuthController::class, 'me']);
    Route::post('/login', [AuthController::class, 'login']);
});

Route::middleware(['throttle:25,1'])->prefix('user')->group(function () {
    Route::post('/register', [AuthController::class, 'register']);
});
