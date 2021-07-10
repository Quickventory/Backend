<?php

namespace App\Console\Commands;

use Illuminate\Foundation\Console\RequestMakeCommand as Command;

class RequestMakeCommand extends Command
{
    protected function getDefaultNamespace($rootNamespace)
    {
        return "{$rootNamespace}\Requests";
    }
}
