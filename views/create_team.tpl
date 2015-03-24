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
					
                		    	<label>Chinese name*</label><br>
                		    	<input name="chinese_name" type="text" value="{{.init.Ch_name}}"/><br>

                		    	<label>English name*</label><br>
                		    	<input name="english_name" type="text" value="{{.init.En_name}}"/><br>

                		    	<label>Member 1 chinses name*</label><br>
                		    	<input name="mem1_chname" type="text" value="{{.init.Mem1chname}}"/><br>
					<span class="help-block">Your Chinese name. Example: 张三丰</span> 
                		    	<label>Member 1 english name*</label><br>
                		    	<input name="mem1_enname" type="text" value="{{.init.Mem1enname}}"/><br>
					<span class="help-block">The phonetic transcription of your Chinese name. Example: ZhangSanfeng (Without blank!!)</span> 

				       	<label class="radio-inline">
				       		<input type="radio"  value="0" name="sex1" {{if compare .init.Sex1 0}} checked="true" {{end}}><label id="radio_label">Male</label>
					</label>
				       	<label class="radio-inline">
				       		<input type="radio"  value="1" name="sex1" {{if compare .init.Sex1 1}} checked="true" {{end}}><label id="radio_label">Female</label>
				       	</label><br>
                		    	<label>Member 1 email*</label><br>
                		    	<input name="mem1_email" type="text" value="{{.init.Mem1email}}"/><br>

                		    	<label>Member 2 chinses name*</label><br>
                		    	<input name="mem2_chname" type="text" value="{{.init.Mem2chname}}"/><br>
                		    	<label>Member 2 english name*</label><br>
                		    	<input name="mem2_enname" type="text" value="{{.init.Mem2enname}}"/><br>

				       	<label class="radio-inline">
				       		<input type="radio"  value="0" name="sex2" {{if compare .init.Sex2 0}} checked="true" {{end}}><label id="radio_label">Male</label>
					</label>
				       	<label class="radio-inline">
				       		<input type="radio"  value="1" name="sex2" {{if compare .init.Sex2 1}} checked="true" {{end}}><label id="radio_label">Female</label>
				       	</label><br>

                		    	<label>Member 2 email*</label><br>
                		    	<input name="mem2_email" type="text" value="{{.init.Mem2email}}"/><br>

                		    	<label>Member 3 chinses name*</label><br>
                		    	<input name="mem3_chname" type="text" value="{{.init.Mem3chname}}"/><br>
                		    	<label>Member 3 english name*</label><br>
                		    	<input name="mem3_enname" type="text" value="{{.init.Mem3enname}}"/><br>
					
				       	<label class="radio-inline">
				       		<input type="radio"  value="0" name="sex3" {{if compare .init.Sex3 0}} checked="true" {{end}}><label id="radio_label">Male</label>
					</label>
				       	<label class="radio-inline">
				       		<input type="radio"  value="1" name="sex3" {{if compare .init.Sex3 1}} checked="true" {{end}}><label id="radio_label">Female</label>
				       	</label><br>

                		    	<label>Member 3 email*</label><br>
                		    	<input name="mem3_email" type="text" value="{{.init.Mem3email}}"/><br>

				    	<input type="hidden" name="uid" value="{{.init.Tid}}">
					<button type="submit" class="btn btn-success">Submit</button>
				</fieldset>
			</form>
		</div>
	</div>
</div>

{{template "tail.tpl" .}}

