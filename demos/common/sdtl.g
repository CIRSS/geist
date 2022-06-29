{{ prefix "sdtl" "https://rdf-vocabulary.ddialliance.org/sdtl#" }}

{{ macro "sdtl_program_node_style" '''
    node[shape=box style="filled" fillcolor="#CCFFCC" peripheries=1 fontname=Courier]
''' }}

{{ query "sdtl_select_program" '''
    SELECT ?program
    WHERE {
        ?program rdf:type sdtl:Program .
    }
'''}}

{{ query "sdtl_select_load_commands" "ProgramID" "VariableName" '''
    SELECT ?command ?dataframe_id ?dataframe_name ?source_line ?source_text ?variable_name
    WHERE {
        <{{$ProgramID}}> sdtl:Commands ?command_inventory .
        ?command_inventory rdfs:member ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:ProducesDataframe ?dataframe_id .
        ?dataframe_id sdtl:VariableInventory/rdfs:member/sdtl:VariableName ?variable_name .
        FILTER (?variable_name="{{$VariableName}}")
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }
'''}}

{{ query "sdtl_select_save_commands" "ProgramID" '''
    SELECT ?command ?dataframe_id ?dataframe_name ?source_line ?source_text
    WHERE {
        <{{$ProgramID}}> sdtl:Commands ?command_inventory .
        ?command_inventory rdfs:member ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:ConsumesDataframe ?dataframe_id .
        ?dataframe_id sdtl:DataframeName ?dataframe_name .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }
'''}}

{{ query "sdtl_select_update_commands" "ProgramID" "VariableName" '''
    SELECT ?command ?source_line ?source_text
    WHERE {
        <{{$ProgramID}}> sdtl:Commands ?command_inventory .
        ?command_inventory rdfs:member ?command .
        ?command rdf:type sdtl:Compute .
        ?command sdtl:Variable ?variable .
        ?variable sdtl:VariableName ?variable_name .
        FILTER (?variable_name="{{$VariableName}}")
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }
'''}}

{{ query "sdtl_select_variable_flows" "ProgramID" "VariableName" '''
    SELECT ?producer_line ?producer_text ?consumer_line ?consumer_text
    WHERE {
        <{{$ProgramID}}> sdtl:Commands ?command_inventory .
        ?command_inventory rdfs:member ?command .
        ?command rdf:type sdtl:Compute .
        ?command sdtl:Variable ?variable .
        ?variable sdtl:VariableName ?variable_name .
        FILTER (?variable_name="{{$VariableName}}")
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?producer_line .
        ?source_info sdtl:OriginalSourceText ?producer_text .
        BIND (?producer_line AS ?consumer_line)
        BIND (?producer_text AS ?consumer_text)
    }
'''}}

{{ query "sdtl_select_dataframe_variables" "DataframeID" '''
    SELECT ?variable
    WHERE {
        <{{$DataframeID}}> sdtl:VariableInventory/rdfs:member/sdtl:VariableName ?variable .
    }
'''}}

{{ query "sdtl_select_commands" "ProgramID" '''
    SELECT DISTINCT ?command ?source_text
    WHERE {
        $ProgramID sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?command .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }
''' }}

{{ query "sdtl_get_command_source" "CommandID" '''
    SELECT DISTINCT ?source_text
    WHERE {
        <{{$CommandID}}> sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }
''' }}

{{ query "sdtl_select_compute_variable_compute_edges" "ProgramID" '''
    SELECT DISTINCT ?upstream_command ?downstream_command ?variable_name
    WHERE {
        $ProgramID sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?downstream_command .
        ?upstream_command sdtl:ProducesDataframe ?dataframe .
        ?downstream_command sdtl:ConsumesDataframe  ?dataframe .
        ?downstream_command sdtl:OperatesOn/sdtl:VariableName ?variable_name .
        ?upstream_command sdtl:Variable/sdtl:VariableName ?variable_name .
    }
''' }}

{{ query "sdtl_select_load_variable_compute_edges" "ProgramID" '''
    SELECT DISTINCT ?upstream_command ?downstream_command ?variable_name
    WHERE {
        $ProgramID sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?downstream_command .
        ?downstream_command sdtl:ConsumesDataframe  ?dataframe .
        ?downstream_command sdtl:OperatesOn/sdtl:VariableName ?variable_name .
        ?upstream_command a sdtl:Load .
        ?upstream_command_load_command sdtl:ProducesDataframe/sdtl:VariableInventory ?variable_name .
    }
''' }}

{{ query "sdtl_select_variable_flow" "ProgramID" "VariableName" '''
    SELECT ?command ?variable
    WHERE {
        <{{$ProgramID}}> sdtl:Commands ?command_inventory .
        ?command_inventory rdfs:member ?command
		{{ consumes_variable "?command" "?variable" }}
    }
'''}}

{{ rule "program_command" "Program" "Command" '''
	{
        {{_subject $Program}} rdf:type sdtl:Program .
        {{_subject $Program}} sdtl:Commands ?_command_inventory .
        ?_command_inventory rdfs:member {{_object $Command}} .
	}
''' }}

{{ rule "variable_writer" "Command" "VariableName" '''
	{
        {
            {{_subject $Command}} sdtl:Variable ?_write_variable .
        }
        UNION
        {
            {{_subject $Command}} rdf:type sdtl:Load .
            {{_subject $Command}} sdtl:ProducesDataframe ?_write_dataframe .
            ?_write_dataframe sdtl:VariableInventory ?_write_variable_inventory .
            ?_write_variable_inventory rdfs:member ?_write_variable .
        }
		?_write_variable sdtl:VariableName {{_object $VariableName}}
	}
''' }}

{{ rule "variable_reader" "Command" "VariableName" '''
	{
        {
            {{_subject $Command}} sdtl:OperatesOn ?_read_variable .
        }
		UNION
		{
            {{_subject $Command}} rdf:type sdtl:Save .
            {{_subject $Command}} sdtl:ConsumesDataframe ?_read_dataframe .
            ?_read_dataframe sdtl:VariableInventory ?_read_variable_inventory .
            ?_read_variable_inventory rdfs:member ?_read_variable .
		}
		?_read_variable sdtl:VariableName {{_object $VariableName}}
	}
''' }}

{{ rule "command_source" "Command" "SourceLine" "SourceText" '''
	{
		{{_subject $Command}} sdtl:SourceInformation ?_source_info .
		?_source_info sdtl:LineNumberStart {{_object $SourceLine}} .
		?_source_info sdtl:OriginalSourceText {{_object $SourceText}}
	}
''' }}

{{ rule "upstream_variable_writer" "Variable" "Command" "Writer" '''
	{
		{{_subject $Command}} (sdtl:ConsumesDataframe/^sdtl:ProducesDataframe)+ {{_object $Writer}}
		{{ variable_writer $Writer $Variable }}
	}
''' }}

{{ rule "intermediate_variable_writer" "Variable" "Writer" "Reader" '''
    {
        {{ upstream_variable_writer $Variable $Reader "?intermediate_writer" }} .
        {{ upstream_variable_writer $Variable "?intermediate_writer" $Writer  }}
    }
''' }}

{{ rule "variable_write_read_edge" "Variable" "Writer" "Reader" '''
    {
        {{ variable_reader $Reader $Variable }}
        {{ upstream_variable_writer $Variable $Reader $Writer }} .
		FILTER NOT EXISTS {
			{{ intermediate_variable_writer $Variable $Writer $Reader }}
		}
    }
''' }}

{{ query "sdtl_select_variable_write_read_edges" "Program" '''
    SELECT DISTINCT ?variable ?writer ?reader
    WHERE {
        {{ program_command $Program "?reader" }} .
        {{ variable_write_read_edge  "?variable" "?writer" "?reader" }}
    } ORDER BY ?variable ?reader ?writer
''' }}

{{ rule "upstream_dataframe_writer" "Reader" "Writer" '''
	{
		{{_subject $Reader}} (sdtl:ConsumesDataframe/^sdtl:ProducesDataframe)+ {{_object $Writer}}
	}
''' }}

{{ query "sdtl_select_dataframe_edges" "ProgramID" '''
    SELECT DISTINCT ?upstream_command ?downstream_command ?dataframe ?dataframe_name
    WHERE {
        $ProgramID sdtl:Commands ?commandinventory .
        ?commandinventory ?index ?upstream_command .
        ?upstream_command sdtl:ProducesDataframe ?dataframe .
        ?downstream_command sdtl:ConsumesDataframe  ?dataframe .
        ?dataframe sdtl:DataframeName ?dataframe_name
    }
''' }}
