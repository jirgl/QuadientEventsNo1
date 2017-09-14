import { getTask } from '../../src/api/endpoints';

describe('API test', () => {

    it('try to find best path', () => {
        getTask().then((data) => {
            console.log(data);
        }).catch((err) => {
            console.log(err);
        });
    });

});
