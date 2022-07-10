# RSS

## About

The RSS system needs to be able generate classic RSS format output.  The issue
of caching may be an issue as the number of resources going into the generation 
of the RSS feed increases.  

It might be easy to cache to the object store but that makes that a requirement
in the architecture.  Perhaps some sort of sqlite or KV store approach could
be used.   Or perhaps a Parquet file based approach falling back again on the 
object store (which we have as a requirement anyway).

The question is do we need zoltan to have some sort of regular queue running 
and updating some of these synication feeds rather than trying to do them on 
the fly on request?

