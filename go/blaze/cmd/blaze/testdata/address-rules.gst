

{{ query "GetEmailForFirstName" "FirstName" '''
    SELECT ?email
    WHERE
    {
        ?person ab:firstname "Craig"^^<http://www.w3.org/2001/XMLSchema#string> .
        ?person ab:email     ?email .
    }
''' }}
