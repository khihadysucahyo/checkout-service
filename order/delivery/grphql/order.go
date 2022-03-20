package grphql

// order schema graphql
const orderSchema = `
	type Order {
		id: ID!
		user: User!
		items: [ItemOrder]!
		total: Float!
		createdAt: String!
	}

	type ItemOrder {
		id: ID!
		item: Item!
		quantity: Int!
		total: Float!
	}

	type Item {
		id: ID!
		name: String!
		price: Float!
		sku: String!
	}

	type User {
		id: ID!
		name: String!
		email: String!
	}

	type Query {
		getOrder(id: ID!): Order
		getOrders: [Order]
	}

	type Mutation {
		createOrder(userId: ID!, items: [ItemOrderInput]!): Order
	}

	input ItemOrderInput {
		itemId: ID!
		quantity: Int!
	}
`
