Thread.abort_on_exception = true

desc "Run locally and clean up when done"
task :server do
  sh "docker build -t kube_leader ."

  begin
    sh "kubectl apply -f test/deployment.yaml"

    loop do
      puts "Waiting for server to start ..."
      sleep 1
      status = `kubectl get pods -l app=kube-leader`
      break if status.scan(/Running/).size == 2

      if status.include?("CrashLoopBackOff")
        abort "Failed to start:\n" + `kubectl logs -l project=kube-leader --previous`
      end
    end

    puts "Streaming logs ... Ctrl+c to shut down ..."
    sh "stern -l app=kube-leader"
  rescue Interrupt
    # user stopped ... do not print a backtrace
  ensure
    sh "kubectl delete -f test/deployment.yaml --force --grace-period 0 --ignore-not-found"
  end
end
