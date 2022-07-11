<?php

header('Content-Type: application/json');

setcookie('backend', $_SERVER['HTTP_HOST'], time() + 3600);

$message = [
    "message" => "Hello world",
    "cookies" => $_COOKIE
];

echo json_encode($message);
