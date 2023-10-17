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
    _id: ObjectId("652e92a807e9476c2c6a9376"),
    email: "login@test.com",
    password: Binary(Buffer.from("JDJhJDEwJGpJeWNhYzFlbWd1azdQT29vM1V3b2VSdWRrRnBsRS80ZXp1cUVFS1RHODViSGRuT004ZzR1","hex"), 0)
  },
  {
    _id: ObjectId("652e92d107e9476c2c6a9377"),
    email: "alreadyIn@use.com",
    password: Binary(Buffer.from("JDJhJDEwJFhhcmtUV1BZWGhZUU41SlJwZlVpVE9TTVIzUGZnT3RvR1VhSHoxOUJUQWdyL0lmRkpXT2JH","hex"), 0)
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
