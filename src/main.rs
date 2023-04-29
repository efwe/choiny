use bson::DateTime;
use futures::stream::TryStreamExt;
use mongodb::{bson::doc, Client, options::ClientOptions};
use mongodb::bson::oid::ObjectId;
use mongodb::options::FindOptions;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct Route {
    _id: ObjectId,
    created_at: DateTime,
    ascend: f64,
    descend: f64,
    content: String,
    title: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct TrackPoint {
    _id: ObjectId,
    route_id: ObjectId,
    time: DateTime,
    location: [f64; 2],
    elevation: f64,
}


#[tokio::main]
async fn main() -> mongodb::error::Result<()> {
    // Replace the placeholder with your Atlas connection string
    let uri = "mongodb://localhost:27017";
    let client_options =
        ClientOptions::parse(uri)
            .await?;
    // Create a new client and connect to the server
    let client = Client::with_options(client_options)?;
    // Send a ping to confirm a successful connection
    client
        .database("admin")
        .run_command(doc! {"ping": 1}, None)
        .await?;
    println!("Pinged your deployment. You successfully connected to MongoDB!");

    // List the names of the collections in that database.
    // for collection_name in db.list_collection_names(None).await? {
    //     println!("{}", collection_name);
    // }
    // Get a handle to a collection of `Book`.
    let db = client.database("wipu");
    let route_collection = db.collection::<Route>("routes");
    let point_collection = db.collection::<TrackPoint>("track_points");
    let filter = doc! { "_id": ObjectId::parse_str("51737e6c2bb882d748000674").unwrap() };
    let find_options = FindOptions::builder().sort(doc! { "track_start": 1 }).build();
    let mut cursor = route_collection.find(filter, find_options).await?;

    // Iterate over the results of the cursor.
    while let Some(route) = cursor.try_next().await? {
        println!("{}", serde_json::to_string_pretty(&route).unwrap());
        let point_filter = doc! { "route_id": route._id };
        let point_find_options = FindOptions::builder().sort(doc! { "time": 1 }).build();
        let mut point_cursor = point_collection.find(point_filter, point_find_options).await?;
        while let Some(point) = point_cursor.try_next().await? {
            println!("{}", serde_json::to_string_pretty(&point).unwrap());
        }
    }
    Ok(())
}