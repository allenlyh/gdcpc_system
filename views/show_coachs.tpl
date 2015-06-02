
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
							<center>School</center>
						</th>
						<th>
							<center>Email</center>
						</th>
						<th>                      
							<center>东苑住宿</center>
						</th>
						<th>                      
							<center>是否愿意拼房</center>
						</th>
					</tr>
				</thead>
				<tbody>
					{{range $key, $val := .coachs}}
						<tr>
							<td>
								<center>{{$key}}</center>
							</td>
							<td>
								<center>{{$val.Chname}}</center>
							</td>
							<td>
								<center>{{$val.School}}</center>
							</td>
							<td>
								<center>{{$val.Email}}</center>
							</td>
							<td>
								<center>{{$val.Accomodate}}</center>
							</td>
							<td>
								<center>{{if compare $val.Shareroom 0}} 愿意 {{else}} 不愿意 {{end}}</center>
							</td>
						</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
