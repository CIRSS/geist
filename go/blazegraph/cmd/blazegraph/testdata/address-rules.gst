

{{ query "GetEmailForFirstName" '''
    SELECT ?email
    WHERE
    {
        ?person ab:firstname "{{.}}"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?person ab:email     ?email .
    }
''' }}