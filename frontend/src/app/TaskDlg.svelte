<script>
    import { getContext } from "svelte";
    import { getapi, postapi, putapi } from "./api.js";
    import Dialog from "./Lib/Dialog.svelte";
    import FormInput from "./Lib/FormInput.svelte";
    import FormText from "./Lib/FormText.svelte";
    import FormCombo from "./Lib/FormCombo.svelte";
    import FormNumber from "./Lib/FormNumber.svelte";
    import FormCheckList from "./Lib/FormCheckList.svelte";
    import FormCheck from "./Lib/FormCheck.svelte";

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
            values = {};
            linked = {};
            update();
            dialog.show();
        },
    };
    let linked = {};
    let options = {};
    let dialog;
    let changed = null;

    const ctx = getContext("ctx");
    let lng = ctx.lng;

    let values = {};
    let req = { date: "Дата", title: "Задача" };

    let repeats = [
        { id: 0, title: "Не повторять" },
        { id: 1, title: "Повторять с интервалом" },
        { id: 2, title: "Еженедельно" },
        { id: 3, title: "Ежемесячно" },
        { id: 4, title: "Ежегодно" },
    ];
    let weeks = [
        { id: 1, title: "пн" },
        { id: 2, title: "вт" },
        { id: 3, title: "ср" },
        { id: 4, title: "чт" },
        { id: 5, title: "пт" },
        { id: 6, title: "сб" },
        { id: 7, title: "вс" },
    ];
    let months = [
        { id: 1, title: "январь" },
        { id: 2, title: "февраль" },
        { id: 3, title: "март" },
        { id: 4, title: "апрель" },
        { id: 5, title: "май" },
        { id: 6, title: "июнь" },
        { id: 7, title: "июль" },
        { id: 8, title: "август" },
        { id: 9, title: "сентябрь" },
        { id: 10, title: "октябрь" },
        { id: 11, title: "ноябрь" },
        { id: 12, title: "декабрь" },
    ];
    let mode = 0;
    let day = 1;
    let dayMonth = "";
    let monthItems = {};
    let weekItems = {};
    let prelast = false;
    let last = false;

    function parseRepeat(repeat) {
        for (let i = 0; i < months.length; i++) {
            monthItems[months[i].id] = true;
        }
        for (let i = 0; i < weeks.length; i++) {
            weekItems[weeks[i].id] = false;
        }
        if (!repeat) {
            mode = 0;
            return;
        }
        let pars = repeat.split(" ");
        switch (repeat[0]) {
            case "d":
                mode = 1;
                if (pars.length > 1) {
                    day = pars[1];
                }
                break;
            case "w":
                mode = 2;
                if (pars.length > 1) {
                    let dweek = pars[1].split(",");
                    for (let i = 0; i < dweek.length; i++) {
                        let d = dweek[i];
                        if (d > 0 && d < 8) {
                            weekItems[d] = true;
                        }
                    }
                }
                break;
            case "m":
                mode = 3;
                if (pars.length > 1) {
                    let dm = pars[1].split(",");
                    let days = [];
                    for (let i = 0; i < dm.length; i++) {
                        let d = dm[i];
                        if (d > 0 && d < 32) {
                            days.push(d);
                        } else if (d == -1) {
                            last = true;
                        } else if (d == -2) {
                            prelast = true;
                        }
                    }
                    dayMonth = days.join(",");
                }
                if (pars.length > 2) {
                    for (let i = 0; i < months.length; i++) {
                        monthItems[months[i].id] = false;
                    }
                    let dm = pars[2].split(",");
                    for (let i = 0; i < dm.length; i++) {
                        let d = dm[i];
                        if (d > 0 && d < 13) {
                            monthItems[d] = true;
                        }
                    }
                }
                break;
            case "y":
                mode = 4;
                break;
            default:
                mode = 0;
        }
    }

    function joinRepeat() {
        let repeat = "";
        switch (mode) {
            case 1:
                if (day <= 0) {
                    day = 1;
                }
                repeat += "d " + day;
                break;
            case 2:
                repeat += "w";
                let dweek = [];
                Object.keys(weekItems).forEach(function (key) {
                    if (weekItems[key]) {
                        dweek.push(key);
                    }
                });
                if (dweek.length > 0) {
                    repeat += " " + dweek.join(",");
                }
                break;
            case 3:
                repeat += "m ";
                let vals = [];
                if (dayMonth) {
                    let tmp = dayMonth.split(",");
                    for (let i = 0; i < tmp.length; i++) {
                        let ti = tmp[i].trim();
                        if (ti >= 1 && ti <= 31) {
                            vals.push(ti);
                        }
                    }
                }
                if (last) {
                    vals.push("-1");
                }
                if (prelast) {
                    vals.push("-2");
                }
                if (vals.length == 0) {
                    repeat += "x";
                } else {
                    repeat += vals.join(",");
                }

                let m = [];
                Object.keys(monthItems).forEach(function (key) {
                    if (monthItems[key]) {
                        m.push(key);
                    }
                });
                if (m.length > 0 && m.length != 12) {
                    repeat += " " + m.join(",");
                }
                break;
            case 4:
                repeat += "y";
                break;
        }
        return repeat;
    }

    function update() {
        if (options.id != "") {
            getapi(
                "task?id=" + options.id,
                (data) => {
                    values = { ...data };
                    let date = values["date"];
                    values["date"] =
                        date.substr(6) +
                        "." +
                        date.substr(4, 2) +
                        "." +
                        date.substr(0, 4);
                    parseRepeat(values["repeat"]);
                },
                ctx.fn.error,
            );
        } else {
            values = {
                date: "",
                title: "",
            };
        }
    }

    function success(data, args) {
        changed = null;
        /*        if (data.message) {
            ctx.fn.info(data.message);
        }*/
        options.callback("update", data, args);
        dialog.close();
    }

    function save() {
        let valpars = {};
        if (options.id != "") {
            /*            Object.keys(changed).forEach(function (key) {
                valpars[key] = values[key];
            });*/
            valpars = { ...values };
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
                (reqs.length == 1 ? lng.ereq : lng.ereqs) + reqs.join(", "),
            );
            return;
        }
        valpars["repeat"] = joinRepeat();
        let d = valpars["date"];
        valpars["date"] = d.substr(6) + d.substr(3, 2) + d.substr(0, 2);
        let args = valpars;
        /*{
            values: valpars,
        };*/
        if (valpars["id"]) {
            putapi(
                "task",
                args,
                (data) => {
                    success(data, args);
                },
                ctx.fn.error,
            );
        } else {
            postapi(
                "task",
                args,
                (data) => {
                    success(data, args);
                },
                ctx.fn.error,
            );
        }
    }

    function changeValue(value) {
        if (!changed) {
            changed = {};
        }
        changed[value] = true;
    }
</script>

<Dialog bind:mod={dialog} header={options.title} icon={options.icon}>
    <div
        slot="content"
        style="padding-bottom: 1em; min-width: 20em;display:flex "
    >
        <div style="padding-right: 1em;">
            <FormInput
                type="date"
                title="Дата"
                width="13"
                req
                change={() => {
                    changeValue("date");
                }}
                bind:value={values["date"]}
            />
            <FormInput
                type="text"
                title="Задача"
                width="20"
                req
                change={() => {
                    changeValue("title");
                }}
                bind:value={values["title"]}
            />
            <FormText
                title="Комментарий"
                size="3"
                change={() => {
                    changeValue("comment");
                }}
                bind:value={values["comment"]}
            />
        </div>
        <div style="padding-left: 1em;">
            <FormCombo
                title="Правило повторения"
                width="20"
                change={() => {
                    changeValue("repeat");
                }}
                list={repeats}
                bind:value={mode}
            />
            {#if mode == 1}
                <FormNumber
                    title="Каждые X дней"
                    bind:value={day}
                    change={() => {
                        changeValue("repeat");
                    }}
                    size="5"
                    width="20"
                />
            {:else if mode == 2}
                <FormCheckList
                    width="25"
                    title="Дни недели"
                    list={weeks}
                    bind:values={weekItems}
                    change={() => {
                        changeValue("repeat");
                    }}
                />
            {:else if mode == 3}
                <FormInput
                    type="text"
                    title="Дни месяца (через запятую)"
                    width="20"
                    req
                    change={() => {
                        changeValue("repeat");
                    }}
                    bind:value={dayMonth}
                />
                <FormCheck
                    title="Предпоследний день месяца"
                    width="20"
                    change={() => {
                        changeValue("repeat");
                    }}
                    bind:value={prelast}
                />
                <FormCheck
                    title="Последний день месяца"
                    width="20"
                    change={() => {
                        changeValue("repeat");
                    }}
                    bind:value={last}
                />
                <FormCheckList
                    width="25"
                    title="Месяцы"
                    list={months}
                    bind:values={monthItems}
                    change={() => {
                        changeValue("month");
                    }}
                />
            {/if}
        </div>
    </div>
    <div
        slot="footer"
        style="flex-grow:1;display:flex;column-gap: 0.5em;justify-content: end;"
    >
        {#if changed}
            <button class="btn primary" style="font-size: 0.9em" on:click={save}
                >Сохранить</button
            >
        {/if}
        <button
            class="btn secondary"
            style="font-size: 0.9em"
            on:click={dialog.close}>{changed ? "Отменить" : "Закрыть"}</button
        >
    </div>
</Dialog>

<style>
</style>
