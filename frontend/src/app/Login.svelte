<script>
    import { setContext } from "svelte";
    import { initTheme, setPath, api } from "./api.js";

    import InputPass from "./Lib/InputPass.svelte";
    import ErrDialog from "./Lib/ErrDialog.svelte";

    let password = "";

    let err;

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

    setContext("ctx", { lng: lng });
    function auth(action) {
        api(
            "signin",
            {
                password: password,
            },
            (data) => {
                if (data.token) {
                    let d = new Date();
                    d.setTime(d.getTime() + 60 * 60 * 8 * 1000); // сутки
                    let expires = "expires=" + d.toUTCString();
                    document.cookie =
                        "token=" + data.token + ";" + expires + ";path=/";
                    window.location = "/";
                }
            },
            (e) => err.show(e),
        );
    }
    function signin() {
        auth("signin");
    }
</script>

<div
    style="min-height: 100vh;display: flex;align-items: center;justify-content: center;overflow:auto"
>
    <div class="card" style="width: 21em;align-items:center;">
        <div
            class="bigger"
            style="width: 100%; display: flex; flex-direction: column; align-items:stretch;padding-bottom: 1em;"
        >
            <div class="smaller">
                <div
                    style="display: flex; align-items: center; justify-content: space-between;margin-bottom: 0.5em"
                >
                    <h5>Введите пароль</h5>
                </div>
            </div>
            <div style="margin: 1em 0;width; 100%">
                <InputPass bind:value={password} placeholder="Пароль" />
            </div>
            <button class="btn primary" on:click={signin}> Войти </button>
        </div>
    </div>
</div>
<ErrDialog bind:err {lng} />

<style>
</style>
