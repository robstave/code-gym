# Answer: Exporting a CSV row

**Difficulty: Medium**

Joining fields with commas does not quote or escape field contents. The input
has two fields, `Ada, Jr.` and `London`, but a CSV reader sees three:
`Ada`, ` Jr.`, and `London`. Fields containing quotes or newlines also need
proper CSV encoding.

Use `encoding/csv.Writer.Write` to encode the record, then call `Flush` and
check `Writer.Error()`. If writing to a buffer, return the buffer's content
only after checking the error.
