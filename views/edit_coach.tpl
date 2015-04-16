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

                		    	<label>Chinese Name*</label><br>
                		    	<input name="chname" type="text" value="{{.init.Chname}}"/><br><br>

                		    	<label>English Name*</label><br>
                		    	<input name="enname" type="text" value="{{.init.Enname}}"/><br><br>

                		    	<label>School Name</label><br>
					<input name="school" type="text" value="{{.init.School}}"/><br><br>

                		    	<label>T-shirt</label><br>
					<input name="tshirt" type="text" value="{{.init.Tshirt}}"/><br><br>

                		    	<label>5月23日晚需要在大学城东苑宾馆住宿名单（包括教练以及队员，以逗号隔开）</label><br>
					<input name="accomodate" type="text" value="{{.init.Accomodate}}"/><br><br>

				    	<input type="hidden" name="uid" value="{{.init.Uid}}">
					<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
