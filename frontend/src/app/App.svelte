<script>
    import { setContext, onMount } from "svelte";
    import { delapi, getapi, postapi } from "./api.js";
    import Alert from "./Lib/Alert.svelte";
    import TaskDlg from "./TaskDlg.svelte";
    import Icon from "./Lib/Icon.svelte";
    import Input from "./Lib/Input.svelte";
    import ErrDialog from "./Lib/ErrDialog.svelte";
    import ConfirmDialog from "./Lib/ConfirmDialog.svelte";
    import InfoDialog from "./Lib/InfoDialog.svelte";

    let err;
    let confirm;
    let info;
    let lng = {
        close: "Закрыть",
        confirm: "Подтверждение",
        ereq: "Заполните поле: ",
        ereqs: "Заполните поля: ",
        error: "Ошибка",
        info: "Информация",
        no: "Нет",
        ok: "ОК",
        qdelete: "Вы действительно хотите удалить задачу?",
        yes: "Да",
    };

    let prevdate = "";
    let fn = {
        info: (m, f) => info.show(m, f),
        error: (e) => {
            if (String(e).includes("401")) {
                window.location = "/login.html";
            } else {
                err.show(e);
            }
        },
        confirm: (m, y, n) => confirm.show(m, y, n),
    };

    setContext("ctx", {
        fn: fn,
        lng: lng,
    });

    let showFormTask;
    let tasks;
    let search = "";
    let istasks = false;

    function update(search) {
        getapi(
            "tasks" + (search ? "?search=" + search : ""),
            (data) => {
                prevdate = "";
                tasks = [];
                istasks = true;
                if (!data.tasks) {
                    return;
                }
                for (let i = 0; i < data.tasks.length; i++) {
                    let item = data.tasks[i];
                    if (prevdate != item.date) {
                        let d =
                            item.date.substr(6) +
                            "." +
                            item.date.substr(4, 2) +
                            "." +
                            item.date.substr(0, 4);
                        tasks.push({
                            date: d,
                            list: [item],
                        });
                        prevdate = item.date;
                    } else {
                        tasks[tasks.length - 1].list.push(item);
                    }
                }
                //                loading = false;
            },
            (e) => {
                if (!String(e).includes("404")) {
                    fn.error(e);
                }
            },
        );
    }

    onMount(() => {
        update();
    });

    function callbackForm(action, param, args) {
        switch (action) {
            case "update":
                update();
                break;
        }
    }

    function formTask(id, opts) {
        let title = id == "" ? "Добавить новую задачу" : "Карточка задачи";
        showFormTask.show({
            title: title,
            param: opts && opts.param ? opts.param : null,
            id: id,
            callback: callbackForm,
        });
    }

    function addtask() {
        formTask("");
    }

    function done(id) {
        postapi(
            "task/done?id=" + id,
            {},
            (data) => {
                update();
            },
            fn.error,
        );
    }

    function edit(item) {
        formTask(item.id);
    }

    function trash(item) {
        fn.confirm(lng.qdelete, () => {
            delapi(
                "task?id=" + item.id,
                {},
                (data) => {
                    update();
                },
                fn.error,
            );
        });
    }
    function searchtasks() {
        update(search);
    }
</script>

<!--svelte:window bind:innerWidth={x} /-->
<div class="app">
    <div class="topnav">
        <button class="btn primary" on:click={addtask}>
            Добавить задачу
        </button>
        <div style="display:flex;column-gap: 0.5em">
            <Input
                style="width: 10em;"
                placeholder="Поиск..."
                bind:value={search}
            />
            <button
                class="btn primary"
                disabled={!search}
                on:click={searchtasks}>Найти</button
            >
            {#if search}
                <button
                    class="btn secondary"
                    on:click={() => {
                        search = "";
                        update();
                    }}>Отмена</button
                >
            {/if}
        </div>
    </div>
    <div class="body">
        <div
            style="flex-grow:1;overflow: auto;flex-direction: column;padding: 1em 2em"
        >
            {#if tasks}
                {#each tasks as day}
                    <div class="day">{day.date}</div>
                    <div class="notelist">
                        {#each day.list as item}
                            <div class="notecard">
                                <!-- svelte-ignore a11y-click-events-have-key-events -->
                                <div
                                    class="panel note"
                                    style="position: relative"
                                >
                                    <div class="todo">
                                        {#if item.repeat}
                                            <div class="fav">
                                                <Icon
                                                    color="var(--gray-500)"
                                                    name="repeat-variant"
                                                    size="1.25em"
                                                />
                                            </div>
                                        {/if}
                                        <button
                                            class="btnicon"
                                            style="padding-bottom: 0.5em"
                                            on:click={() => {
                                                done(item.id);
                                            }}
                                        >
                                            <!-- name="checkbox-blank-circle-outline"
                                name="checkbox-marked-circle-outline" -->
                                            <svg
                                                class="tocheck"
                                                viewBox="0 0 24 24"
                                            >
                                                <path
                                                    d="M12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"
                                                />
                                            </svg>
                                        </button>
                                        <div class="notetitle">
                                            {item.title}
                                        </div>
                                    </div>
                                    {#if item.comment}<div class="notetext">
                                            {item.comment}
                                        </div>{/if}
                                    <!--                                {#if item.project}
                                    <div class="project">
                                        {item.project}
                                    </div-->
                                    <div class="notebtns">
                                        <a
                                            href={null}
                                            title="Выполнить"
                                            on:click={() => {
                                                done(item.id);
                                            }}
                                            ><Icon
                                                name="check-bold"
                                                size="1.4em;"
                                            /></a
                                        >
                                        <a
                                            href={null}
                                            title="Редактировать"
                                            on:click={() => {
                                                edit(item);
                                            }}
                                            ><Icon
                                                name="edit"
                                                size="1.4em;"
                                            /></a
                                        >
                                        <a
                                            href={null}
                                            title="Удалить"
                                            on:click={(e) => {
                                                trash(item);
                                            }}
                                            ><Icon
                                                name="trash-can-outline"
                                            /></a
                                        >
                                    </div>
                                </div>
                            </div>
                        {/each}
                    </div>
                {/each}
                {#if tasks.length == 0 && istasks}
                    <div style="display: flex; justify-content: center">
                        <Alert
                            type="warning"
                            msg="Задачи не найдены"
                            show
                            noclose
                        />
                    </div>
                {/if}
            {/if}
        </div>
    </div>
</div>
<TaskDlg bind:showForm={showFormTask} />

<!--ContextMenu {event} bind:showMenu={showContext} items={contextItems} /-->
<ErrDialog bind:err {lng} />
<ConfirmDialog bind:confirm {lng} />
<InfoDialog bind:info {lng} />

<!--SelectInTable bind:show={showSelect} /-->

<style>
    .app {
        height: 100vh;
        display: flex;
        flex-direction: column;
    }
    .body {
        flex-grow: 1;
        display: flex;
        flex-direction: column;
        min-height: 0;
        position: relative;
        /*      	flex: 1;*/
    }

    .topnav {
        background-color: var(--cardbg-color);
        border-bottom: var(--border-width) solid var(--card-border-color);
        top: 0;
        width: 100%;
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        padding: 0.5em 1em;
        column-gap: 1em;
    }
    :global(.notelist) {
        margin: 1em 0;
        columns: 20em;
        /*        flex-grow: 1;*/
    }

    :global(.notecard) {
        padding-bottom: 1em;
        break-inside: avoid;
    }

    :global(.note) {
        position: relative;
        cursor: default;
        font-size: 0.9em;
        padding: 0.5em 1em;
        /*margin-bottom: 1em;*/
        break-inside: avoid;
    }
    /*    .note:first-child {
        margin-top: 0;
    }*/
    :global(.notetitle) {
        font-weight: 600;
        padding-bottom: 0.5em;
    }

    :global(.notebtns) {
        display: flex;
        align-items: center;
        justify-content: right;
        column-gap: 0.5em;
        visibility: hidden;
        fill: var(--gray-700);
    }

    :global(.note:hover .notebtns) {
        visibility: visible;
    }
    :global(.fav) {
        position: absolute;
        top: 0.5em;
        right: 0.5em;
    }
    .day {
        display: flex;
        align-items: center;
        column-gap: 0.5em;
        font-size: 1.2em;
        font-weight: 600;
        padding: 0.25em 0em;
        border-bottom: 2px dotted var(--gray-500);
    }
    .tocheck {
        width: 1.5em;
        height: 1.5em;
        fill: var(--font-color);
    }
    .tocheck:hover {
        fill: var(--primary);
    }

    .tocheck:hover path {
        d: path(
            "M20,12A8,8 0 0,1 12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4C12.76,4 13.5,4.11 14.2,4.31L15.77,2.74C14.61,2.26 13.34,2 12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12M7.91,10.08L6.5,11.5L11,16L21,6L19.59,4.58L11,13.17L7.91,10.08Z"
        );
        d: "M20,12A8,8 0 0,1 12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4C12.76,4 13.5,4.11 14.2,4.31L15.77,2.74C14.61,2.26 13.34,2 12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12M7.91,10.08L6.5,11.5L11,16L21,6L19.59,4.58L11,13.17L7.91,10.08Z";
    }
    .todo {
        display: flex;
        align-items: center;
        column-gap: 0.4em;
    }
</style>
