# Simple Elastic Search Bulk Importer

Simple Elastic Bulk importer. This importer leverage the ability of Microsoft SQL Server to convert output to json files. The workflow is as follows, a stored procedure is generating json's and the Go API is retrieving these jsons and making sure the the order of these jsons are correctly set.

## Requirements
- Elasticsearch 
- Microsoft SQL Server 2016

## Usage app
```
# run dep ensure 
dep ensure

# Go to root and run or build an executable
go run ./cmd/api/main.go

```




Below an example query that generates the correct JSON 

```
 -- Modify your query here, make sure you give a DocID / BulkInstructions and Source (This is the actual data)
    SELECT cast(DATEDIFF(s, '1970-01-01 00:00:00.000', d.[Date] ) as bigint) DocID, -- Primary key
          -- >>>>>>>>>>>>>>>>>>>>>> Adjust the index and add the id, modify column names to <index_name>._id &  <index_name>._index
          (SELECT @Index 'create._index',
                  cast(DATEDIFF(s, '1970-01-01 00:00:00.000', d.[Date] ) as bigint) 'create._id'
          FROM dbo.DateDimension DD 
          WHERE DD.[Date] = D.[Date] -- Primary key 
          FOR JSON PATH, WITHOUT_ARRAY_WRAPPER) BulkInstructions,
          -- >>>>>>>>>>>>>>>>>>>>>> Source query, fill in the query to populate your source in Elasticsearch
          (SELECT *
          FROM dbo.DateDimension DS
            WHERE DS.[Date] = D.[Date]  
            FOR JSON AUTO , WITHOUT_ARRAY_WRAPPER  ) Source
          -- >>>>>>>>>>>>>>>>>>>>>>  
    FROM dbo.DateDimension D
    ORDER BY D.[Date]
    -- >>>>>>>>>>>>>>>>>>>>>> Make sure that you are not importing everything at ones 
    OFFSET @Offset ROWS 
    FETCH NEXT @Fetch ROWS ONLY
```

To use it go to the following api

http://localhost:8888/v1/bulk?db_name=...&db_user=...&db_pass=...&elastic_index=...&fetch=...&offset=...

## Author
Sten

