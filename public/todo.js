(function($) {
'use strict';
$(function() {
var todoListItem = $('.todo-list');
var todoListInput = $('.todo-list-input');
$('.todo-list-add-btn').on("click", function(event) {
	event.preventDefault();

	var item = $(this).prevAll('.todo-list-input').val();

	if (item) {
		//서버에서 json object로 반환해주기 때문에 값을 그냥 사용가능하다
		$.post("todos/",{name:item},additem)
		//function(e){
		//addItem({name : item, completed : false});
	//});
//todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
	todoListInput.val("");
	}

});

todoListItem.on('change', '.checkbox', function() {
	var id = $(this).closest("li").attr('id');
	var $self = $(this);
	var compeleted = true;
	if ($self.attr('checked')){
		completed = false;
	}
	$.get("complete-todo/"+id+"?complete="+complete,function(data){
		
		if (complete) {
			$self.removeAttr('checked');
		} else {
			$self.attr('checked', 'checked');
		}

		$self.closest("li").toggleClass('completed');
	});
});

var addItem = function(item){
	if (item.completed){
		todoListItem.append("<li class='completed'"+"id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
	
	}else{
		todoListItem.append("<li" + "'id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");

	}
};

$.get('/todos',function(items){
	items.forEach(e => {
		addItem(e)
	});
});

todoListItem.on('click', '.remove', function() {
	//url : todos/id method : DELETE
	var id = $(this).closest("li").attr('id');
	var $self = $(this);
	$.ajax({
		url : "todos/" + id,
		type : "DELETE",
		success : function(data){
			if (data.success){
				$self.parent().remove();
			}
		}
	});
	//$(this).parent().remove();
});

});
})(jQuery);
