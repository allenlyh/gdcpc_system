{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<table class="table table-striped table-hover table-bordered">
				<thead>
					<tr>
						<th>
							<center>No.</center>
						</th>
						<th>
							<center>Coach Name</center>
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
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .total_teams}}
						<tr>
							<td>
								<center>{{$key}}</center>
							</td>
							<td>
								<center>{{$val.Coachname}}</center>
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
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
