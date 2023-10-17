/* global use, db */
// Select the echo-api database
use("echo-api");
//User collection
db.createCollection("user", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      title: "User Object Validation",
      required: ["email", "password"],
      properties: {
        email: {
          bsonType: "string",
          description: "'email' must be a string and is required",
        },
        password: {
          bsonType: "binData",
          description: "'password' encrypted using bcrypt",
        },
      },
    },
  },
});

//Book collection
db.createCollection("book", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      title: "Book Object Validation",
      required: ["name", "owner_id"],
      properties: {
        name: {
          bsonType: "string",
          description: "Book name",
        },
        owner_id: {
          bsonType: "objectId",
          description: "User ID",
        },
      },
    },
  },
});
// Insert a few documents into the sales collection.
db.getCollection('user').insertMany([
  {
    _id: ObjectId("651f2d851990790417cce077"),
    email: 'login@test.com',
    password: Binary(Buffer.from("243261243130242e67446743436d3245695343437341424e6d67735775434361474243652f4764756173656871735343507a30675a315353674f6975", "hex"), 0)
  }
]);

// Run a find command to view items sold on April 4th, 2014.
// const salesOnApril4th = db.getCollection('sales').find({
//   date: { $gte: new Date('2014-04-04'), $lt: new Date('2014-04-05') }
// }).count();

// Print a message to the output window.
// console.log(`${salesOnApril4th} sales occurred in 2014.`);

// Here we run an aggregation and open a cursor to the results.
// Use '.toArray()' to exhaust the cursor to return the whole result set.
// You can use '.hasNext()/.next()' to iterate through the cursor page by page.
// db.getCollection('sales').aggregate([
//   // Find all of the sales that occurred in 2014.
//   { $match: { date: { $gte: new Date('2014-01-01'), $lt: new Date('2015-01-01') } } },
//   // Group the total sales for each product.
//   { $group: { _id: '$item', totalSaleAmount: { $sum: { $multiply: [ '$price', '$quantity' ] } } } }
// ]);
