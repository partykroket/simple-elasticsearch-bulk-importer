/****** Object:  StoredProcedure [Api].[OUTPUT_Update_Option_Json_DEV]    Script Date: 15/01/2019 18:06:11 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

/* 
==========================================================================================================
Author		Sten Raadschelders
Project        Elasticsearch Data Inserter
Description 	Preparation, staging of data.

USAGE:

- To generate data:

   EXEC [dbo].[GEN_Elastic_Bulk] @Index = 'index_name', @BulkType = 'create', @Offset = 0 , @Fetch = 250

WORKFLOW:
  1. 
  

Updates
[2018-06-12] - Created...
==========================================================================================================
*/

ALTER PROCEDURE [dbo].[GEN_Elastic_Bulk] @Index nvarchar(50) ,
                                              @BulkType nvarchar(25) = 'create',
                                              @Offset int, 
                                              @Fetch int 
AS
BEGIN

    -- SET NOCOUNT ON added to prevent extra result sets from
    -- interfering with SELECT statements.
    SET NOCOUNT ON;

    -- * 1. Initialize

    -- >>>>>>>>>>>>>>>>>>>>>> Create bulk table
	  CREATE TABLE #BulkResultData ([DocID] int,
										[BulkInstructions] nvarchar(max),
										[Source] nvarchar(max))
    -- >>>>>>>>>>>>>>>>>>>>>> Create bulk table

    -- >>>>>>>>>>>>>>>>>>>>>> Start populate query!
    INSERT INTO #BulkResultData ( [DocID],
                                 [BulkInstructions],
                                [Source])
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
    -- >>>>>>>>>>>>>>>>>>>>>> Make sure that you are not importing everything at ones 

    -- >>>>>>>>>>>>>>>>>>>>>> Get results 
    select * 
    from #BulkResultData
    -- >>>>>>>>>>>>>>>>>>>>>> Get results 

END