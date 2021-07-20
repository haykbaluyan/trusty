package fcc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var htmlTest string = `





<html>

<head>
	<title>FCC Registration System</title>
	<link rel="stylesheet" type="text/css" href="images/registration/cores.css">
</head>

<body>








								      
     
  
            
        
            
  
  
 
  							      
    
            
        
           
  

 

<div id="srchDetailContent">


<div id="srchDetailClose"><a href="javascript:self.close()">Close Window</a></div>


<div id="srchDetailHeader">Registration Detail</div>       

<table>

<tr>
	<th>FRN:</th>
	<td>0000010827</td>
</tr>

<tr>
	<th>Registration Date:</th>
	<td>09/15/2000 08:30:58 AM</td>
</tr>

<tr>
	<th>Last Updated:</th>
	<td>
	          	
		04/05/2021 12:12:43 PM
	</td>
</tr>

<tr>
	<th>Business Name:</th>
	<td>	
	
		Veracity Networks, LLC
	
	</td>
</tr>

<!-- <##dba##> -->

<tr>
	<th>Business Type:</th>
	<td>
	          	
		Private Sector
										
	
	,
										
	          	
		Limited Liability Corporation
	
	</td>
</tr>

<tr>
	<th>Contact Organization:</th>
	<td>
	          	
		
										
	</td>
</tr>

<tr>
	<th>Contact Position:</th>
	<td>
	          	
		FCC Contact
										
	</td>
</tr>
<tr>
	<th>Contact Name:</th>
	<td>

	
	    Tara Lyle 
	

	</td>
</tr>

<tr>
	<th>Contact Address:</th>
	<td>
	  
	
	
357 S. 670 W.<br>
       

Ste 300<br>
       

Lindon, UT 84042<br>
       
       

United States                                      
                                            

	
	</td>
</tr>

<tr>
	<th>Contact Email:</th>
	<td>
	          	
		tara.lyle@veracitynetworks.com
	
	</td>
</tr>

<tr>
	<th>ContactPhone:</th>
	<td >
	          	
		(801) 878-3225 
	
	</td>
</tr>

<tr>
	<th>ContactFax:</th>
	<td>
	          	
		(801) 373-0682
										
	</td>
</tr>

</table>

</div> <!-- Close srchDetailContent -->


						
						


</body>
</html>
`

func TestNewSearchDetailFromHTML(t *testing.T) {
	sd := NewSearchDetailFromHTML(htmlTest)
	email, err := sd.GetEmail()
	require.NoError(t, err)
	assert.Equal(t, "tara.lyle@veracitynetworks.com", email)
}