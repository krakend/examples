<?php

declare(strict_types=1);

header('Content-Type: application/json');

function getUserData(string $username, array $users): array {
    return $users[$username] ?? [];
}

function getUserReviews(int $userId, array $reviews): array {
    return $reviews[(string)$userId] ?? [];
}

$users = [
    'admin' => [
        'user_id' => 2,
        'user_name' => 'admin',
    ],
];

$reviews = [
    '2' => [
        [
            'review_id' => 1,
            'review_title' => 'First example',
        ],
        [
            'review_id' => 2,
            'review_title' => 'Second example',
        ],
    ],
];

$message = [];

try {
    $path = $_SERVER['PATH_INFO'] ?? '';
    switch ($path) {
        case '/user':
            if (empty($_POST['username'])) {
                throw new InvalidArgumentException('Username not provided', 400);
            }

            $message['user_data'] = getUserData($_POST['username'], $users);
            if (empty($message['user_data'])) {
                throw new RuntimeException('User not found', 404);
            }
            break;

        case '/reviews':
            if (empty($_POST['user_id'])) {
                throw new InvalidArgumentException('User ID not provided', 400);
            }

            $userId = filter_var($_POST['user_id'], FILTER_VALIDATE_INT);
            if ($userId === false) {
                throw new InvalidArgumentException('Invalid user ID', 400);
            }

            $message['reviews'] = getUserReviews($userId, $reviews);
            break;

        default:
            throw new RuntimeException('Invalid endpoint', 404);
    }

    echo json_encode($message, JSON_THROW_ON_ERROR);
} catch (InvalidArgumentException $e) {
//    http_response_code($e->getCode());
    echo json_encode(['error' => $e->getMessage()], JSON_THROW_ON_ERROR);
} catch (RuntimeException $e) {
//    http_response_code($e->getCode());
    echo json_encode(['error' => $e->getMessage()], JSON_THROW_ON_ERROR);
} catch (JsonException $e) {
//    http_response_code(500);
    echo json_encode(['error' => 'JSON encoding error'], JSON_THROW_ON_ERROR);
}
