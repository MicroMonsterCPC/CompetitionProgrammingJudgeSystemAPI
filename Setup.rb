require './judge'

class Setup
  def main
    lang = select_lang
    user_name = input_user_name
    js = JudgeSystem.new(user_name, lang); js.main
  end
  def select_lang
    puts "使用した言語を入力してください"
    lang = {
      "ruby"    => "rb",
      "crystal" => "cr",
      "c#"      => "cs"
    }[gets.chomp]
    return lang
  end
  def input_user_name
    puts "あなたのGitHubUserNameを入力してください"
    return gets.chomp
  end
end

Setup.new.main
