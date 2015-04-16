{{template "head.tpl" .}}

<div class="container theme-showcase" role="main">
	<div class="row">
		<div class="span12">
			<legend>珠海赛区</legend> 
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
							<center>School</center>
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
							<center>T-shirt</center>
						</th>
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .zh_teams}}
						<tr>
							<td>
								<center>{{$key}}</center>
							</td>
							<td>
								<center>{{$val.Coachnamech}}</center>
								<center>{{$val.Coachnameen}}</center>
								<center>{{$val.Coachemail}}</center>
							</td>
							<td>
								<center>{{$val.School}}</center>
							</td>
							<td>
								<center>{{$val.Ch_name}}</center>
								<center>{{$val.En_name}}</center>
							</td>
							<td>
							<B>
								<center><font {{if compare $val.Sex1 1}} color="red" {{end}}>{{$val.Mem1chname}}</font></center>
								<center><font {{if compare $val.Sex1 1}} color="red" {{end}}>{{$val.Mem1enname}}</font></center>
								<center><font {{if compare $val.Sex1 1}} color="red" {{end}}>{{$val.Mem1email}}</font></center>
							</B>
							</td>
							<td>
							<B>
								<center><font {{if compare $val.Sex2 1}} color="red" {{end}}>{{$val.Mem2chname}}</font></center>
								<center><font {{if compare $val.Sex2 1}} color="red" {{end}}>{{$val.Mem2enname}}</font></center>
								<center><font {{if compare $val.Sex2 1}} color="red" {{end}}>{{$val.Mem2email}}</font></center>
							</B>
							</td>
							<td>
							<B>
								<center><font {{if compare $val.Sex3 1}} color="red" {{end}}>{{$val.Mem3chname}}</font></center>
								<center><font {{if compare $val.Sex3 1}} color="red" {{end}}>{{$val.Mem3enname}}</font></center>
								<center><font {{if compare $val.Sex3 1}} color="red" {{end}}>{{$val.Mem3email}}</font></center>
							</B>
							</td>
							<td>
								<center>{{$val.Tshirt}}</center>
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>

			<legend>广州赛区</legend> 
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
							<center>School</center>
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
							<center>T-shirt</center>
						</th>
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .gz_teams}}
						<tr>
							<td>
								<center>{{$key}}</center>
							</td>
							<td>
								<center>{{$val.Coachnamech}}</center>
								<center>{{$val.Coachnameen}}</center>
								<center>{{$val.Coachemail}}</center>
							</td>
							<td>
								<center>{{$val.School}}</center>
							</td>
							<td>
								<center>{{$val.Ch_name}}</center>
								<center>{{$val.En_name}}</center>
							</td>
							<td>
							<B>
								<center><font {{if compare $val.Sex1 1}} color="red" {{end}}>{{$val.Mem1chname}}</font></center>
								<center><font {{if compare $val.Sex1 1}} color="red" {{end}}>{{$val.Mem1enname}}</font></center>
								<center><font {{if compare $val.Sex1 1}} color="red" {{end}}>{{$val.Mem1email}}</font></center>
							</B>
							</td>
							<td>
							<B>
								<center><font {{if compare $val.Sex2 1}} color="red" {{end}}>{{$val.Mem2chname}}</font></center>
								<center><font {{if compare $val.Sex2 1}} color="red" {{end}}>{{$val.Mem2enname}}</font></center>
								<center><font {{if compare $val.Sex2 1}} color="red" {{end}}>{{$val.Mem2email}}</font></center>
							</B>
							</td>
							<td>
							<B>
								<center><font {{if compare $val.Sex3 1}} color="red" {{end}}>{{$val.Mem3chname}}</font></center>
								<center><font {{if compare $val.Sex3 1}} color="red" {{end}}>{{$val.Mem3enname}}</font></center>
								<center><font {{if compare $val.Sex3 1}} color="red" {{end}}>{{$val.Mem3email}}</font></center>
							</B>
							</td>
							<td>
								<center>{{$val.Tshirt}}</center>
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
