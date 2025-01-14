<?php

header('Content-Type: application/json');

$message = [
    "message" => "Hello world"
];

if ('true' === ($_GET['return_error'] ?? ''))
{
    header($_SERVER["SERVER_PROTOCOL"] . ' 500 Internal Server Error', true, 500);
    $message = [
        "error" => "There was an error: " . $_SERVER["SERVER_PROTOCOL"] . ' 500 Internal Server Error'
    ];
}

echo json_encode($message);
