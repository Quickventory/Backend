<?php

namespace App\Console\Commands;

use Illuminate\Routing\Console\ControllerMakeCommand as Command;

class ControllerMakeCommand extends Command
{
    protected function getDefaultNamespace($rootNamespace)
    {
        return "{$rootNamespace}\Controllers";
    }
}
