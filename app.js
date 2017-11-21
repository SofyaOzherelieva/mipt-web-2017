var tag_ul = document.createElement('ul');

set('Сделать задание #3 по web-программированию');

function del() {
    this.parentNode.remove();
}


document.getElementById('root').appendChild(tag_ul);

var input = document.createElement('input');
input.id = 'add_task_input';
var button1 = document.createElement('button');
button1.id = 'add_task';
button1.innerHTML = 'Добавить';

document.getElementById('root').appendChild(input);
document.getElementById('root').appendChild(button1);


function set(string) {
    var tag_li = document.createElement('li');
    var tag_span = document.createElement('span');
    var button = document.createElement('button');
    tag_span.innerHTML = string;
    button.innerHTML = 'Удалить';
    button.addEventListener('click', del);

    tag_ul.appendChild(tag_li);
    tag_li.appendChild(tag_span);
    tag_li.appendChild(button);
}

button1.addEventListener('click', function () {
    set(input.value);
});