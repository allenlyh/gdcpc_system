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
					
                		    	<label>Username</label>
					<p> {{.show.Username}} </p>

                		    	<label>Chinese Name</label><br>
					<p> {{.show.Chname}} </p>

                		    	<label>English Name</label><br>
					<p> {{.show.Enname}} </p>

                		    	<label>Email</label><br>
					<p> {{.show.Email}} </P>

                		    	<label>School Name</label><br>
					<p> {{.show.School}} </p>

                		    	<label>T-shirt</label><br>
					<p> {{.show.Tshirt}} </p>

                		    	<label>5月23日晚需要在大学城东苑宾馆住宿名单（包括教练以及队员）</label><br>
					<p> {{.show.Accomodate}} </p>

					<label>{{if compare .show.Shareroom 0}} 愿意 {{else}} 不愿意 {{end}} 拼房 </label>
				</fieldset>
			</form>
<!--			<button type="button" class="btn btn-success" onclick="window.location='/edit_coach'">Setting</button><br> -->
			<legend>珠海赛区</legend> 
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
							<center>T-shirt</center>
						</th>
						<th>                      
							<center>Action</center>
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
								<br>
								<center>{{$val.Tshirt}}</center>
							</td>
							<td>
								<button type="button" class="btn btn-primary" onclick="window.location='/create_team?tid={{$val.Tid}}'">Edit</button>
								<button type="button" class="btn btn-danger" onclick="window.location='/action?action=DeleteTeam&tid={{$val.Tid}}'">Delete</button>
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
						<th>                      
							<center>Action</center>
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
								<br>
								<center>{{$val.Tshirt}}</center>
							</td>
							<td>
								<button type="button" class="btn btn-primary" onclick="window.location='/create_team?tid={{$val.Tid}}'">Edit</button>
								<button type="button" class="btn btn-danger" onclick="window.location='/action?action=DeleteTeam&tid={{$val.Tid}}'">Delete</button>
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
