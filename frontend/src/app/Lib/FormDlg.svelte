<script>
    import { getContext } from "svelte";
    import { module } from "../api.js";
    import Dialog from "./Dialog.svelte";
    import Form from "../Lib/Form.svelte";

    export const showForm = {
        show(opts) {
            changed = null;
            options = opts;
            if (!options.icon) {
                options.icon = {};
            }
            if (!options.param) {
                options.param = {};
            }
            ctrls = [];
            values = {};
            files = {};
            linked = {};
            update();
            dialog.show();
        },
    };
    let files = {};
    let linked = {};
    let options = {};
    let dialog;
    let changed = null;

    const ctx = getContext("ctx");
    let lng = ctx.lng;

    let table = "";
    let ctrls = [];
    let values = {};
    let calc = {};
    let req = {};

    function update() {
        let obj = { url: "api/dbform", param: {} };
        if (options.api) {
            let api = options.api("dbform");
            if (api) {
                obj = api;
            }
        }
        module(
            obj.url,
            {
                name: options.form,
                id: options.id,
                ...obj.param,
                ...options.param,
            },
            (data) => {
                if (options.clone) {
                    options.id = "";
                }
                table = data.table;
                ctrls = data.ctrls;
                values = data.values;
                if (data.calc) {
                    calc = data.calc;
                } else {
                    calc = {};
                }
                req = {};
                getReq(ctrls);
            },
            ctx.fn.error
        );
    }

    function success(data, args) {
        changed = null;
        //                    edit = 0;
        files = {};
        /*  if (id != "") {
                            getRow();
                        }*/
        if (data.message) {
            ctx.fn.info(data.message);
        }
        // tablefunc("update", data.id);
        options.callback("update", data, args);
        dialog.close();
    }

    function getReq(ctrls) {
        for (let i = 0; i < ctrls.length; i++) {
            let ctrl = ctrls[i];
            if (ctrl.req) {
                req[ctrl.field] = ctrl.title;
            }
            if (ctrl.ctrls) {
                getReq(ctrl.ctrls);
            }
        }
    }

    function replaceCtrl(ctrls, ctrl) {
        for (let i = 0; i < ctrls.length; i++) {
            if (ctrls[i].field == ctrl.field) {
                ctrls[i] = ctrl;
                return false;
            }
            if (ctrl.ctrls) {
                let ret = replaceCtrl(ctrl.ctrls, ctrl);
                if (!ret) {
                    return false;
                }
            }
        }
        return true;
    }

    function save() {
        let valpars = {};
        if (options.id != "") {
            Object.keys(changed).forEach(function (key) {
                valpars[key] = values[key];
            });
        } else {
            valpars = { ...values };
        }
        Object.keys(linked).forEach(function (key) {
            valpars[key] = linked[key];
        });
        let reqs = [];
        for (const key in valpars) {
            if (valpars.hasOwnProperty(key)) {
                if (!valpars[key] && req[key]) {
                    reqs.push(req[key]);
                }
            }
        }
        if (reqs.length > 0) {
            ctx.fn.error(
                (reqs.length == 1 ? lng.ereq : lng.ereqs) + reqs.join(", ")
            );
            return;
        }
        let args = {
            table: table,
            id: options.id, //edit == 2 ? "" : id,
            values: valpars,
        };
        let upload = false;
        for (const [key, value] of Object.entries(files)) {
            if (value.length > 0) {
                upload = true;
                break;
            }
        }
        let obj = { url: "api/dbupdate", param: {} };
        if (options.api) {
            let api = options.api("dbupdate");
            if (api) {
                obj = api;
            }
        }
        if (upload) {
            moduleform(
                obj.url,
                files,
                args,
                (data) => {
                    success(data, args);
                },
                ctx.fn.error
            );
        } else {
            module(
                obj.url,
                args,
                (data) => {
                    success(data, args);
                },
                ctx.fn.error
            );
        }
    }

    function callback(action, value) {
        switch (action) {
            case "change":
                if (!changed) {
                    changed = {};
                }
                changed[value] = true;
                if (calc[value]) {
                    module(
                        calc[value],
                        { ctrl: value, value: values[value] },
                        (data) => {
                            for (let i = 0; i < data.ctrls.length; i++) {
                                replaceCtrl(ctrls, data.ctrls[i]);
                            }
                            ctrls = ctrls;
                        },
                        ctx.fn.error
                    );
                }
                //options.callback("change", value, values[value]);
                break;
        }
    }
</script>

<Dialog bind:mod={dialog} header={options.title} icon={options.icon}>
    <div slot="content" style="padding-bottom: 1em; min-width: 20em;">
        <Form {ctrls} bind:values bind:linked {callback} />
    </div>
    <div
        slot="footer"
        style="flex-grow:1;display:flex;column-gap: 0.5em;justify-content: end;"
    >
        {#if changed}
            <button class="btn primary" style="font-size: 0.9em" on:click={save}
                >{lng.save}</button
            >
        {/if}
        <button
            class="btn secondary"
            style="font-size: 0.9em"
            on:click={dialog.close}>{changed ? lng.cancel : lng.close}</button
        >
    </div>
</Dialog>

<style>
</style>
