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
					
                		    	<label>username*</label><br>
                		    	<input name="username" type="text"/> 
                		    	<span class="help-block">Only letter and digit are allowed.</span> 
                		    	<span class="help-block">3-20 characters</span><br>

                		    	<label>Password*</label><br>
					<input name="password" type="password"/>
                		    	<span class="help-block">6-20 characters</span><br>

                		    	<label>Password Confirm*</label><br>
                		    	<input name="password2" type="password"/><br><br>

                		    	<label>Name*</label><br>
                		    	<input name="name" type="text"/><br><br>

                		    	<label>Email*</label><br>
					<input name="email" type="text"/><br><br>

                		    	<label>School Name*</label><br>
					<input name="school" type="text"/><br><br>

				    	<input type="hidden" name="uid" value="{{.init.Uid}}">
					<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}
