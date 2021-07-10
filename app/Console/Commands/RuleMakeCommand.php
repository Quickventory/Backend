<?php

namespace App\Console\Commands;

use Illuminate\Foundation\Console\RuleMakeCommand as Command;

class RuleMakeCommand extends Command
{
    protected function getDefaultNamespace($rootNamespace)
    {
        return "{$rootNamespace}\Rules";
    }
}
