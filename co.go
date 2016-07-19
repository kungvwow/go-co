//
// Package co : async/await for glolang
//
package co

//
// Task definition
//
type Task struct {
	// communicate via Channel
	Channel chan interface{}

	// store the result of the Task
	Result interface{}

	// store error
	Error error
}

//
// Async : create a new *Task
//
// execute fn , save the result, send to channel
//
// e.g
// co.Async(func() interface{}{
// 	return val
// 	// val will be Task's Result
// })
//
func Async(fn func() interface{}) *Task {
	t := new(Task)
	t.Channel = make(chan interface{})

	// run the task
	// collect the result
	// set as the ret Task's Channel
	go func() {
		// final work
		defer func() {
			if err := recover(); err != nil {
				t.Error = err.(error)
			}

			// error is a finish state too
			t.Channel <- t.Result
		}()

		t.Result = fn()
	}()

	return t
}

//
// Await : await a Task & return it's result
//
// e.g
// res := co.Await(Task)
//
func Await(t *Task) (interface{}, error) {
	// when t.Channel is available
	// set result as await ret value
	// `result = await(Task)`
	t.Result = <-t.Channel

	return t.Result, t.Error
}