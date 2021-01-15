
{{ macro "OriginalSourceText" '''
        {{ "sdtl:OriginalSourceText" }}
'''}}

{{ query "sdtl_get_command_source" "CommandID" '''
    SELECT DISTINCT ?source_text
    WHERE {
        <{{$CommandID}}> sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }
''' }}
