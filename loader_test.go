package ini

import (
    "testing"
)

func TestRemoveComments( t *testing.T ) {
    s := "logfile=/var/log/supervisor/supervisord.log \\; ; (main log file;default $CWD/supervisord.log)"
    if removeComments( s ) != "logfile=/var/log/supervisor/supervisord.log \\;" {
        t.Fail()
    }
}
