<script>
    import BtnIcon from "./BtnIcon.svelte";
    // https://svelte.dev/repl/f391b3186f804fb0bb0931e73388553a?version=3.46.4

    export let value = ruiso(new Date());
    /*    export let days = "Su|Mo|Tu|We|Th|Fr|Sa".split("|");
    export let months = "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec".split(
        "|"
    );*/
    export let days = "пн|вт|ср|чт|пт|сб|вс".split("|");
    export let months =
        "январь|февраль|март|апрель|май|июнь|июль|август|сентябрь|октябрь|ноябрь|декабрь".split(
            "|",
        );
    export let start = 1; // first day of the week (0 = Sunday, 1 = Monday)
    export let offset = 0; // offset in months from currently selected date
    export let callback;

    let date = iso(new Date());
    let today = iso(new Date());

    function ru2iso(value) {
        let pars = value.split(".");
        switch (pars.length) {
            case 1:
                return value;
            case 2:
                return pars[1] + "-" + pars[0];
        }
        return pars[2] + "-" + pars[1] + "-" + pars[0];
    }

    function iso2ru(date) {
        let pars = date.split("-");
        switch (pars.length) {
            case 1:
                return value;
            case 2:
                return pars[1] + "." + pars[2];
        }
        return pars[2] + "." + pars[1] + "." + pars[0];
    }

    function ruiso(date) {
        const pad = (n) => (n < 10 ? "0" + n : n);
        return (
            pad(date.getDate()) +
            "." +
            pad(date.getMonth() + 1) +
            "." +
            date.getFullYear()
        );
    }

    function iso(date) {
        const pad = (n) => (n < 10 ? "0" + n : n);
        return (
            date.getFullYear() +
            "-" +
            pad(date.getMonth() + 1) +
            "-" +
            pad(date.getDate())
        );
    }

    $: acceptDate(value);

    function acceptDate(value) {
        const newDate = new Date(ru2iso(value));
        offset = 0;
        if (!isNaN(newDate)) {
            date = iso(newDate);
        }
    }

    function go(direction) {
        offset = offset + direction;
    }

    function selectDate(newValue) {
        date = newValue;
        value = iso2ru(newValue);
        offset = 0;
        if (callback) {
            callback();
        }
    }

    $: viewDate = viewDateFrom(date, offset);

    $: month = months[viewDate.getMonth()];

    $: year = viewDate.getFullYear();

    $: weeks = weeksFrom(viewDate, date, start);

    function viewDateFrom(date, offset) {
        var viewDate = new Date(date);
        viewDate.setDate(1);
        viewDate.setMonth(viewDate.getMonth() + offset);
        return viewDate;
    }

    function weeksFrom(viewDate, date, start) {
        var first = new Date(viewDate.getTime());
        first.setDate(1);
        first.setDate(first.getDate() + ((start - first.getDay() - 7) % 7));

        var last = new Date(viewDate.getTime());
        last.setDate(
            new Date(
                viewDate.getFullYear(),
                viewDate.getMonth() + 1,
                0,
            ).getDate(),
        );
        last.setDate(last.getDate() + ((start - last.getDay() + 6) % 7));

        var d = new Date(first.getTime()),
            M = viewDate.getMonth(),
            Y = viewDate.getFullYear(),
            week = [],
            weeks = [week];

        while (d.getTime() <= last.getTime()) {
            var dd = d.getDate(),
                mm = d.getMonth(),
                yy = d.getFullYear(),
                value = iso(d);

            week.push({
                date: dd,
                value,
                class: [
                    date === value
                        ? "dateselected"
                        : today == value
                          ? "today"
                          : "",
                    mm == M
                        ? ""
                        : (mm > M ? yy >= Y : yy > Y)
                          ? "future"
                          : "past",
                ].join(" "),
            });

            d = new Date(d.getFullYear(), d.getMonth(), d.getDate() + 1);

            if (d.getDay() === start) {
                week = [];
                weeks.push(week);
            }
        }

        return weeks;
    }
</script>

<div>
    <table class="datepicker">
        <tr>
            <td style="vertical-align: middle;line-height: 0">
                <BtnIcon
                    name="angle-left"
                    fn={() => {
                        go(-1);
                    }}
                />
            </td>
            <td colspan="5">{month} {year}</td>
            <td style="vertical-align: middle;line-height: 0">
                <BtnIcon
                    name="angle-right"
                    fn={() => {
                        go(1);
                    }}
                />
            </td>
        </tr>
        <tr>
            {#each days as day}
                <th>{day}</th>
            {/each}
        </tr>
        {#each weeks as week}
            <tr>
                {#each week as day}
                    <!-- svelte-ignore a11y-click-events-have-key-events -->
                    <td
                        class="datebtn {day.class}"
                        on:click={() => selectDate(day.value)}>{day.date}</td
                    >
                {/each}
            </tr>
        {/each}
    </table>
</div>

<style>
    .datepicker td,
    .datepicker th {
        width: 2em;
        text-align: center;
        border-radius: 0.25em;
        line-height: 1.5em;
        margin: 0;
        padding: 0;
    }
    td.past,
    td.future {
        opacity: 0.5;
    }
    .datebtn {
        cursor: pointer;
    }
    .datebtn:hover {
        background: var(--gray-300);
        color: var(--font-color);
    }
    .dateselected {
        color: var(--pbtn-color);
        font-weight: 600;
        background-color: var(--pbtn-bg);
        border-color: var(--pbtn-bg);
    }
    .today {
        color: var(--sbtn-color);
        font-weight: 600;
        background-color: var(--sbtn-bg);
        border-color: var(--sbtn-bg);
    }
</style>
