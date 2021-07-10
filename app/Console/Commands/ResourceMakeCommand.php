<?php

namespace App\Console\Commands;

use Illuminate\Foundation\Console\ResourceMakeCommand as Command;

class ResourceMakeCommand extends Command
{
    protected function getDefaultNamespace($rootNamespace)
    {
        return "{$rootNamespace}\Resources";
    }
}
