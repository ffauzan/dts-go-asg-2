# Assignment 2

## How to Run

### 1. Setting Up the Environment

Copy the content of `.env.template` to `.env` and fill in all the required environment variables. Alternatively, you can set environment variables directly on the system.

### 2. Applying Migrations

Apply migrations from the `migrations` folder.

### 3. Starting the Application

Start the app. The program entry point is located in `cmd/server/main.go`.


## REST API Endpoints

### 1. `POST /orders`

#### Request Body (application/json)

```json
{
    "customerName": "Customer 1",
	"items": [
		{
			"itemCode": "C1ITEM1",
			"description": "C1ITEM1 desc",
			"quantity": 4
		},
        {
			"itemCode": "C1ITEM2",
			"description": "C1ITEM2 desc",
			"quantity": 2
		}
	]
}
```

### 2. `GET /orders`

#### Response Body (application/json)

```json
{
	"status": 200,
	"message": "Orders",
	"data": [
        {
            "id": 1,
            "customerName": "Customer 1",
            "items": [
                {
                    "itemCode": "C1ITEM1",
                    "description": "C1ITEM1 desc",
                    "quantity": 4
                },
                {
                    "itemCode": "C1ITEM2",
                    "description": "C1ITEM2 desc",
                    "quantity": 2
                }
            ]
        },
         {
            "id": 2,
            "customerName": "Customer 2",
            "items": [
                {
                    "itemCode": "C2ITEM1",
                    "description": "C1ITEM1 desc",
                    "quantity": 4
                },
                {
                    "itemCode": "C1ITEM2",
                    "description": "C1ITEM2 desc",
                    "quantity": 2
                }
            ]
        }
    ]
}
```

### 3. `DELETE /orders/{id}`

#### Response Body (application/json)

```json
{
    "status": 200,
    "message": "Order deleted",
    "data": null
}
```

### 4. `PUT /orders/{id}`

#### Request Body (application/json)

```json
{
    "customerName": "Customer 1",
    "items": [
        {
            "itemCode": "C1ITEM1",
            "description": "C1ITEM1 desc",
            "quantity": 4
        },
        {
            "itemCode": "C1ITEM2",
            "description": "C1ITEM2 desc",
            "quantity": 2
        },
        {
            "itemCode": "C1ITEM4",
            "description": "C1ITEM4 desc",
            "quantity": 8
        }
    ]
}
```
