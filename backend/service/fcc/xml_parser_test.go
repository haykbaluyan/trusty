package fcc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var xmlForTest string = `
<?xml version="1.0" encoding="ISO-8859-1"?>
<Filer499QueryResults xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="http://apps.fcc.gov/cgb/form499/XMLSchema/Filer499QueryResults_Schema.xsd"
 Updated="2021-06-21" RecordCount="1" >
		
        <Filer>
            <Form_499_ID>831188</Form_499_ID>
            <Filer_ID_Info>
                <Registration_Current_as_of>2021-04-01</Registration_Current_as_of>
                <start_date>2015-01-12</start_date>
                <USF_Contributor>Yes</USF_Contributor>
                <Legal_Name>LOW LATENCY COMMUNICATIONS LLC</Legal_Name>
                <Principal_Communications_Type>Interconnected VoIP</Principal_Communications_Type>
                <holding_company>IPIFONY SYSTEMS INC.</holding_company>
                <FRN>0024926677</FRN>
                <hq_address>
                    <address_line>241 APPLEGATE TRACE</address_line>
                    <city>PELHAM</city>
                    <state>AL</state>
                    <zip_code>35124</zip_code>
                </hq_address>
                <customer_inquiries_address>
                    <address_line>241 APPLEGATE TRACE</address_line>
                    <city>PELHAM</city>
                    <state>AL</state>
                    <zip_code>35124</zip_code>
                </customer_inquiries_address>
                <Customer_Inquiries_telephone>2057453970</Customer_Inquiries_telephone>
                <other_trade_name>Low Latency Communications</other_trade_name>
                <other_trade_name>String by Low Latency</other_trade_name>
                <other_trade_name>Lilac by Low Latency</other_trade_name>
            </Filer_ID_Info>
            <Agent_for_Service_Of_Process>
                <dc_agent>Jonathan Allen Rini O&apos;Neil, PC</dc_agent>
                <dc_agent_telephone>2029553933</dc_agent_telephone>
                <dc_agent_fax>2022962014</dc_agent_fax>
                <dc_agent_email>jallen@rinioneil.com</dc_agent_email>
                <dc_agent_address>
                    <address_line>1200 New Hampshire Ave, NW</address_line>
                    <address_line>Suite 600</address_line>
                    <city>Washington</city>
                    <state>DC</state>
                    <zip_code>20036</zip_code>
                </dc_agent_address>
            </Agent_for_Service_Of_Process>
            <FCC_Registration_information>
                <Chief_Executive_Officer>Daryl Russo</Chief_Executive_Officer>
                <Chairman_or_Senior_Officer>Matthew Hardeman</Chairman_or_Senior_Officer>
                <President_or_Senior_Officer>Larry Smith</President_or_Senior_Officer>
            </FCC_Registration_information>
            <jurisdiction_state>alabama</jurisdiction_state>
            <jurisdiction_state>florida</jurisdiction_state>
            <jurisdiction_state>georgia</jurisdiction_state>
            <jurisdiction_state>illinois</jurisdiction_state>
            <jurisdiction_state>louisiana</jurisdiction_state>
            <jurisdiction_state>north_carolina</jurisdiction_state>
            <jurisdiction_state>pennsylvania</jurisdiction_state>
            <jurisdiction_state>tennessee</jurisdiction_state>
            <jurisdiction_state>texas</jurisdiction_state>
            <jurisdiction_state>virginia</jurisdiction_state>
        </Filer>
</Filer499QueryResults>
`

func TestNewFiler499QueryResultsFromXML(t *testing.T) {
	fq := NewFiler499QueryResultsFromXML(xmlForTest)
	frn, err := fq.GetFRN()
	require.NoError(t, err)
	assert.Equal(t, "0024926677", frn)
}
