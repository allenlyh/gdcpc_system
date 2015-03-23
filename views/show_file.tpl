{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<form id="coach" method="post" action="#">
				<fieldset>
				    <legend>Infomation</legend> 

					{{ if $.warning }}
						<div class="alert alert-warning" role="alert">
							<strong></strong>
							{{.warning}}
						</div>
					{{ end }}
					
					<button type="button" class="btn btn-success" onclick="window.location='/edit_coach'">Edit</button><br>
                		    	<label>Username</label>
					<p> {{.show.Username}} </p>

                		    	<label>Name</label><br>
					<p> {{.show.Name}} </p>

                		    	<label>Email</label><br>
					<p> {{.show.Email}} </P>

                		    	<label>School Name</label><br>
					<p> {{.show.School}} </p>
				</fieldset>
			</form>
			<button type="button" class="btn btn-success" onclick="window.location='/create_team'">Create Team</button><br>
			<table class="table table-striped table-hover table-bordered">
				<thead>
					<tr>
						<th>
							<center>No.</center>
						</th>
						<th>
							<center>Team Name</center>
						</th>
						<th>
							<center>Contestant1</center>
						</th>                    
						<th>                    
							<center>Contestant2</center>
						</th>                      
						<th>                      
							<center>Contestant3</center>
						</th>
						<th>                      
							<center>Action</center>
						</th>
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .total_teams}}
						<tr>
							<td>
								<center>{{$key}}</center>
							</td>
							<td>
								<center>{{$val.Ch_name}}</center>
								<center>{{$val.En_name}}</center>
							</td>
							<td>
								<center>{{$val.Mem1chname}}</center>
								<center>{{$val.Mem1enname}}</center>
								<center>{{$val.Mem1email}}</center>
							</td>
							<td>
								<center>{{$val.Mem2chname}}</center>
								<center>{{$val.Mem2enname}}</center>
								<center>{{$val.Mem2email}}</center>
							</td>
							<td>
								<center>{{$val.Mem3chname}}</center>
								<center>{{$val.Mem3enname}}</center>
								<center>{{$val.Mem3email}}</center>
							</td>
							<td>
								<button type="button" class="btn btn-success" onclick="window.location='/create_team?tid={{$val.Tid}}'">Edit</button><br>
								<button type="button" class="btn btn-success" onclick="window.location='/action?action=DeleteTeam&tid={{$val.Tid}}'">Delete</button><br>
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
