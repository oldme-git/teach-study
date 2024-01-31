<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Facades\DB;

class Controller extends BaseController
{
    use AuthorizesRequests, ValidatesRequests;

    public function mongoTest()
    {
//        $res = DB::table("user")->where("name", "oldme")->first();
//        $res = DB::table("user")->find("65b86cdf583500006f002a12");
        $res = DB::table("user")->insert([
            "name" => "gf",
            "age"   => 3
        ]);
        return $res;
    }
}
