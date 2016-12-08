require './judge'

class Setup
  def main
    lang           = select_lang
    user_name      = input_user_name
    question_id    = input_question_id
    question_level = input_question_level
    js = JudgeSystem.new(user_name, lang, question_id, question_level); js.main
  end

  def select_lang
    puts "使用した言語を入力してください"
    lang = {
      "ruby"       => "rb",
      "crystal"    => "cr",
      "c#"         => "cs",
      "python"     => "py",
      "javascript" => "js"
    }[gets.chomp]
    return lang
  end

  def input_user_name
    puts "あなたのGitHubUserNameを入力してください"
    return gets.chomp
  end

  def input_question_id
    puts "ジャッジするQuestionIdを入力してください(20161125)"
    return gets.to_i
  end

  def input_question_level
    puts "QuestionLevalを入力してください(EXPERT)"
    return gets.chomp
  end
end

Setup.new.main
