var todoApp = angular.module('todoApp', []);
todoApp.controller('TodoListController', function TodoListController ($scope) {
    $scope.todos = [
        {
            name: 'Item 1',
            snippet: "Read Manager's Path book"
        },

    ];
});