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
					
					<label>Username</label><br>
					<p>{{.init.Username}}</p>
					
                		    	<label>Old Password*</label><br>
					<input name="old_password" type="password"/><br><br>

                		    	<label>New Password</label><br>
                		    	<input name="new_password1" type="password"/>
                		    	<span class="help-block">6-20 characters</span>
                		    	<span class="help-block">Ignore this if you don't want to change your password</span><br>

                		    	<label>New Password Confirm</label><br>
                		    	<input name="new_password2" type="password"/><br><br>

                		    	<label>Name</label><br>
                		    	<input name="name" type="text" value="{{.init.Name}}"/><br><br>

                		    	<label>School Name</label><br>
					<input name="school" type="text" value="{{.init.School}}"/><br><br>

				    	<input type="hidden" name="uid" value="{{.init.Uid}}">
					<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
