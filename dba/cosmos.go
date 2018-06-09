package dba

type Item struct {
	Price float64
}

func PriceCheck(itemId int) (float64, bool) {
	item := LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}

func LoadItem(id int) *Item {
	return &Item{
		Price: 9.001,
	}
}

//func QueryToplevelFieldsExamples(t *testing.T, db *mongo.Database) {
//	_, err := db.RunCommand(
//		context.Background(),
//		bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
//	)
//	require.NoError(t, err)
//
//	coll := db.Collection("inventory")
//
//	{
//		// Start Example 6
//
//		docs := []interface{}{
//			bson.NewDocument(
//				bson.EC.String("item", "journal"),
//				bson.EC.Int32("qty", 25),
//				bson.EC.SubDocumentFromElements("size",
//					bson.EC.Int32("h", 14),
//					bson.EC.Int32("w", 21),
//					bson.EC.String("uom", "cm"),
//				),
//				bson.EC.String("status", "A"),
//			),
//			bson.NewDocument(
//				bson.EC.String("item", "notebook"),
//				bson.EC.Int32("qty", 50),
//				bson.EC.SubDocumentFromElements("size",
//					bson.EC.Double("h", 8.5),
//					bson.EC.Int32("w", 11),
//					bson.EC.String("uom", "in"),
//				),
//				bson.EC.String("status", "A"),
//			),
//			bson.NewDocument(
//				bson.EC.String("item", "paper"),
//				bson.EC.Int32("qty", 100),
//				bson.EC.SubDocumentFromElements("size",
//					bson.EC.Double("h", 8.5),
//					bson.EC.Int32("w", 11),
//					bson.EC.String("uom", "in"),
//				),
//				bson.EC.String("status", "D"),
//			),
//			bson.NewDocument(
//				bson.EC.String("item", "planner"),
//				bson.EC.Int32("qty", 75),
//				bson.EC.SubDocumentFromElements("size",
//					bson.EC.Double("h", 22.85),
//					bson.EC.Int32("w", 30),
//					bson.EC.String("uom", "cm"),
//				),
//				bson.EC.String("status", "D"),
//			),
//			bson.NewDocument(
//				bson.EC.String("item", "postcard"),
//				bson.EC.Int32("qty", 45),
//				bson.EC.SubDocumentFromElements("size",
//					bson.EC.Int32("h", 10),
//					bson.EC.Double("w", 15.25),
//					bson.EC.String("uom", "cm"),
//				),
//				bson.EC.String("status", "A"),
//			),
//		}
//
//		result, err := coll.InsertMany(context.Background(), docs)
//
//		// End Example 6
//
//		require.NoError(t, err)
//		require.Len(t, result.InsertedIDs, 5)
//	}
//}
