import { observable } from 'bobx';
import { getTask, ITask, ITaskResult, submitTask } from '../api/endpoints';

class Store {
    @observable private _isAdmin: boolean;

    get isAdmin(): boolean {
        return this._isAdmin;
    }

    acquireTask() {
        getTask().then((task: ITask) => {
            console.log(task);

            // calculate

            submitTask(0, '').then((taskResult: ITaskResult) => {
                console.log(taskResult);
            });
        });
    }
}

export const store = new Store();