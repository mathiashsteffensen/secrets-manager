task default: %w[fmt lint test]

task :test do
    sh "go test ./... --cover"
end

task :fmt do
    sh "gofmt -l -s -w ./"
end

task :lint do
    sh "staticcheck ./..."
end

namespace :version do
    %w[major minor patch].each do |bump_type|
        task bump_type do
            currentVersion = `cat config/version.go | grep VERSION`.split("=")[1].gsub("\"", "").strip

            removeVersionFile(currentVersion)

            index = 0

            if bump_type.to_sym == :minor
                index = 1
            elsif bump_type.to_sym == :patch
                index = 2
            end

            newVersion = currentVersion
                .split(".")
                .map(&:to_i)
                .each_with_index.map { |n, i| i == index ? n.next : n }
                .join(".")

            newVersionFile(newVersion)

            publishVersion(newVersion)
        end
    end
    



    def newVersionFile(version)
        puts "--- Creating version.go file for version #{version}"

        newContents = "package config\n\nconst VERSION = \"#{version}\""

         sh "echo \"#{newContents}\" >> config/version.go"
    end

    def removeVersionFile(version)
        puts "--- Removing version.go file for version #{version}"
        sh "rm config/version.go"
    end

    def publishVersion(version)
        puts "--- Pushing new version to Github"
        sh "
            git add ./config/version.go
            && git commit -m \"bump version to #{version}\"
            && git tag v#{version}
            && git push origin v#{version}
        "
        puts "--- Version #{version} successfully published"
    end
end
