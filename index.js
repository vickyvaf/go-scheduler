document.addEventListener('DOMContentLoaded', function () {
    var quill = new Quill('#editor', {
        theme: 'snow',
        placeholder: 'Tulis konten email di sini...',
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline'],
                [{ 'list': 'ordered' }, { 'list': 'bullet' }],
                ['link', 'clean']
            ]
        }
    });

    // Lucide Icons Initialization
    lucide.createIcons();

    // Initialize Flatpickr with Shadcn Look
    const scheduledInput = flatpickr("#datepicker-trigger", {
        enableTime: true,
        dateFormat: "Z",
        time_24hr: true,
        minDate: "today",
        static: false,
        onOpen: function () {
            document.getElementById('datepicker-trigger').classList.add('ring-2', 'ring-ring', 'ring-offset-2');
        },
        onClose: function (selectedDates, dateStr, instance) {
            document.getElementById('datepicker-trigger').classList.remove('ring-2', 'ring-ring', 'ring-offset-2');
            if (selectedDates.length > 0) {
                // Update the visible label
                const formattedDate = instance.formatDate(selectedDates[0], "d-m-Y H:i");
                document.getElementById('date-label').innerText = formattedDate;
                document.getElementById('date-label').classList.remove('text-muted-foreground');

                // Update the hidden input value
                document.getElementById('scheduled_at').value = dateStr;
            } else {
                document.getElementById('date-label').innerText = "Pilih tanggal dan waktu...";
                document.getElementById('date-label').classList.add('text-muted-foreground');
                document.getElementById('scheduled_at').value = "";
            }
        }
    });

    quill.on('text-change', function () {
        const html = quill.root.innerHTML;
        if (quill.getText().trim().length === 0) {
            document.getElementById('html').value = "";
        } else {
            document.getElementById('html').value = html;
        }
    });

    document.body.addEventListener('htmx:beforeRequest', function () {
        document.getElementById('submitBtn').disabled = true;
        document.getElementById('response').classList.add('hidden');

        const html = quill.root.innerHTML;
        if (quill.getText().trim().length === 0) {
            document.getElementById('html').value = "";
        } else {
            document.getElementById('html').value = html;
        }
    });

    document.body.addEventListener('htmx:afterRequest', function (evt) {
        document.getElementById('submitBtn').disabled = false;
        const resp = document.getElementById('response');
        const alert = document.getElementById('alert-container');
        resp.classList.remove('hidden');

        if (evt.detail.successful) {
            alert.className = "relative w-full rounded-lg border border-green-200 bg-green-50 p-4 text-green-800";
            document.getElementById('emailForm').reset();
            quill.setContents([]);
            // Reset Datepicker Label
            document.getElementById('date-label').innerText = "Pilih tanggal dan waktu...";
            document.getElementById('date-label').classList.add('text-muted-foreground');
        } else {
            alert.className = "relative w-full rounded-lg border border-red-200 bg-red-50 p-4 text-red-800";
            document.getElementById('response-msg').innerText = "Error: " + (evt.detail.xhr.responseText || "Unknown error");
        }
    });
});

