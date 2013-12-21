function bindUpload( browseButton, inputField, progressField, previewField, previewImg, maxSize) {
	var uploader = new plupload.Uploader(
			{
				runtimes : 'html5,html4',
				browse_button : browseButton,
				unique_names : true,
				max_file_size : maxSize,
				url : '/upload',
			});

	uploader.init();
	uploader.bind('FilesAdded', function(up, files) {
		for (var i=0; i<files.length; i++) {
			var fileArr = files[i].name.split(".");
			var fileType = fileArr[fileArr.length-1];
			if (browseButton == "pickFile") {
				if (fileType != "txt") {
					alert("请上传文本文件！");
					uploader.splice();
					return;
				}
			} 
		};
		up.refresh(); // Reposition Flash/Silverlight
		var inputValue = $("input[name='" + inputField + "']").val();
		if (inputValue != undefined) {
			if (inputValue.search("apk")==-1) {
				$("input[name='" + inputField + "']").val("");
			}
		}
		uploader.start();

	});

	uploader.bind('UploadProgress', function(up, file) {
		$('#' + progressField).html(file.percent + "%");
	});

	uploader.bind('Error', function(up, err) {
		$('#' + progressField).html(
				"<div>Error: "
						+ err.code
						+ ", Message: "
						+ err.message
						+ (err.file ? ", File: "
								+ err.file.name : "")
						+ "</div>");
		up.refresh(); // Reposition Flash/Silverlight

	});

	uploader.bind('FileUploaded', function(up, file, response) {
		if(response.response.search("packageName")==-1) {
            $('#' + progressField).html("<span style='color:red;font-weight:bold'>上传失败，请检查网络或者重新登录！</span>");
            return;
        }
		var json = eval("(" + response.response.replace(new RegExp("(<PRE>)|(</PRE>)","g"),"") + ")");
		if (response.response.search("packageName")!=-1) {
			$("label[for='" + inputField + "']").hide("slow");
			$('#' + progressField).html("");
			
			var showStatus = json.itemStatus;
			var auditStatus = json.auditStatus;
			if(showStatus != "SHOW" || auditStatus != "PASS"){
				$("#apkFiles").html('<div><ul class="thumbnails"><li><label class="error" generated="true"><b>'+json.auditDetail+'</b></label></li></ul></div>');
				return;
			}
			$("#apkFiles").html('<div><ul class="thumbnails"><li><a href="javascript:void(0)"><span onclick="deleteApk(event)">删除</span></a></li>'
					+'<li>PackageName<input type="text" value="'+json.packageName+
					'" class="input-medium" disabled/></li><li>VersionCode<input type="text" value="'+json.versionCode+'" class="input-mini" disabled/></li><li>VersionName<input type="text" value="'+json.versionName
					+'" class="input-mini" disabled/></li></ul></div>');
			//$("input[name='" + previewField + "']").val(json.preview);
			$("#uploadPackageName").val(json.packageName);
			$("#uploadMd5").val(json.md5);
			$("#oldVersionName").val(json.versionName);
			$("#oldVersionCode").val(json.versionCode);
			$("label[for='" + inputField + "']").hide("slow");
			$("label[for='" + previewField + "']").hide("slow");
			$('#' + progressField).html("");
			$("#usePublishApk").attr("checked",false);
			//$("input[name='" + previewField + "']").val(json.preview);
			//$("#" + previewImg).attr("src", json.preview);

			return;
		}
		$("input[name='" + inputField + "']").val(json.name);
		$("label[for='" + inputField + "']").hide("slow");
		$("label[for='" + previewField + "']").hide("slow");
		$('#' + progressField).html("");
		//$("input[name='" + previewField + "']").val(json.preview);
		//$("#" + previewImg).attr("src", json.preview);
	});
	
	if($("#" + inputField).val()) {
	}
	if($("input[name='" + previewField + "']").val()) {
		$("#" + previewImg).attr("src", $("input[name='" + previewField + "']").val());
	}

}

function deleteApk(event) {
	if(confirm("确认删除?")) {
		$(event.target).parent().parent().parent().parent().remove();
		if ($("#apkFiles").children().length == 1) {
			if ($("#apkFiles div:first-child").find("input[name='apkDescriptions']").length!=0) {
				$("#apkFiles div:first-child").find("input[name='apkDescriptions']").remove();
				$("#apkFiles div:first-child").find("label[for='apkDescriptions']").remove();
			}
		}
		if ($("#uploadPackageName")){
			$("#uploadPackageName").val("");
		}
		if ($("#uploadMd5")){
			$("#uploadMd5").val("");
		}
		$("#usesPermission").hide();
	}
}

function bindUploadImage( browseButton, inputField, progressField, previewField, previewImg, maxSize) {
	var uploader = new plupload.Uploader(
			{
				runtimes : 'html5,html4',
				browse_button : browseButton,
				unique_names : true,
				max_file_size : maxSize,
				url : '/apps/dev/image/upload',
			});

	uploader.init();
	uploader.bind('FilesAdded', function(up, files) {
		for (var i=0; i<files.length; i++) {
			var fileArr = files[i].name.split(".");
			var fileType = fileArr[fileArr.length-1];
			if (fileType != "jpg" && fileType != "gif" && fileType != "png" && fileType != "jpeg" && fileType != 'bmp' &&
				fileType != "JPG" && fileType != "GIF" && fileType != "PNG" && fileType != "JPEG" && fileType != 'BMP') {
				alert("请上传图片文件！");
				uploader.splice();
				return;
			}
		};
		up.refresh(); // Reposition Flash/Silverlight
		var inputValue = $("input[name='" + inputField + "']").val();
		if (inputValue != undefined) {
			if (inputValue.search("apk")==-1) {
				$("input[name='" + inputField + "']").val("");
			}
		}
		uploader.start();

	});

	uploader.bind('UploadProgress', function(up, file) {
		$('#' + progressField).html(file.percent + "%");
	});

	uploader.bind('Error', function(up, err) {
		$('#' + progressField).html(
				"<div>Error: "
						+ err.code
						+ ", Message: "
						+ err.message
						+ (err.file ? ", File: "
								+ err.file.name : "")
						+ "</div>");
		up.refresh(); // Reposition Flash/Silverlight

	});

	uploader.bind('FileUploaded', function(up, file, response) {
		if(response.response.search("preview")==-1) {
            $('#' + progressField).html("<span style='color:red;font-weight:bold'>上传失败，请检查网络或者重新登录！</span>");
            return;
        }
		var json = eval("(" + response.response.replace(new RegExp("(<PRE>)|(</PRE>)","g"),"") + ")");
		if (response.response.search("preview") ==-1) {
			//$("label[for='" + inputField + "']").hide("slow");
			$('#' + progressField).html("");
			
			$("label[for='" + inputField + "']").hide("slow");
			$("label[for='" + previewField + "']").hide("slow");
			$("#" + previewImg).attr("src", json.preview);

			return;
		}
		$("input[name='" + inputField + "']").val(json.name);
		$("label[for='" + inputField + "']").hide("slow");
		$("label[for='" + previewField + "']").hide("slow");
		$('#' + previewField).val(json.key);
		$("#" + previewImg).attr("src", json.preview);
	});
	if($("input[name='" + previewField + "']").val()) {
		//$("#" + previewImg).attr("src", $("input[name='" + previewField + "']").val());
	}

}
