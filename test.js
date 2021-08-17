const axios = require('/home/chenyc/.nvm/versions/node/v12.20.1/lib/node_modules/axios');

const actionBodyMap = new Map([
    ['GetArticles', { keyword: 'test', condition: { types: [1, 2] , is_deleteds: [1]} }],
    [
        'CreateArticles',
        {
            creations: [
                { type: 1, status: 0, title: '标题1', content: '哈哈哈哈', tag_ids: [1] },
                { type: 2, status: 0, title: '标题2', content: '哈哈哈哈', tag_ids: [1, 2] },
            ],
        },
    ],
    ['UpdateArticles', { updations: [{ id: 1, status: 1, title: 'test', is_deleted: 1, tag_ids: [1, 2, 3] }] }],
    ['DeleteArticles', { ids: [12] }],
    ['GetTags', { condition: {}, keyword: '' }],
    [
        'CreateTags',
        {
            creations: [
                { name: 'TAG1', description: '第9个' },
                { name: 'TAG2', description: '第10个' },
            ],
        },
    ],
    [
        'UpdateTags',
        {
            updations: [
                { id: 1, description: '第1个' },
                { id: 2, description: '第2个' },
            ],
        },
    ],
    ['DeleteTags', { condition: { ids: [3, 4] }, keyword: '' }],
]);

const run = async (action) => {
    const body = actionBodyMap.get(action);
    if (!body) return;
    axios
        .post('http://cc:9527?Action=' + action, body)
        .then((res) => console.log(JSON.stringify(res.data)))
        .catch((err) => console.log(JSON.stringify(err.response.data)));
};

const action = process.argv[2] || 'GetArticles';
run(action);
