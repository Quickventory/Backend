<?php

namespace App\Traits\Controllers;


trait ControllerTrait
{
    public function okResponse($response = null)
    {
        return response()->json($response ?? "", 200);
    }

    public function createdResponse($response = null)
    {
        return response()->json($response ?? "", 201);
    }

    public function unauthorizedResponse($response = null)
    {
        return response()->json($response ?? "", 401);
    }

    public function acceptedResponse($response = null)
    {
        return response()->json($response ?? "", 202);
    }

    public function conflictResponse($response = null)
    {
        return response()->json($response ?? "", 409);
    }

    /**
     * @param string|null $response The response you want to be displayed to the client.
     * @return JsonResponse
     */
    public function deletedResponse($response = null)
    {
        return response()->json($response ?? ['message' => __('general.responses.deleted')], 200);
    }

    public function errorResponse($response = null)
    {
        return response()->json($response ?? ["message" => __('general.responses.error_occurred')], 500);
    }

    public function notFoundResponse($response = null)
    {
        return response()->json($response ?? "", 404);
    }
}
